# TLS Minimum Version

## Summary

The SDK's production client transports require **TLS 1.3** as the minimum
protocol version. This keeps post-quantum (ML-KEM) key exchange available, since
ML-KEM is only negotiated on TLS 1.3. When the floor is left at TLS 1.2 (Go's
default when `MinVersion` is unset), a connection can silently drop below the
level where post-quantum protection is available.

For callers behind legacy middleboxes or proxies that cannot negotiate TLS 1.3,
a configuration option lowers the floor back to **TLS 1.2**. This is an opt-out:
the secure default (1.3) applies unless you explicitly enable it. The option only
lowers the floor — a TLS 1.3-capable server can still negotiate 1.3, keeping
ML-KEM available when both peers support it — but any connection that ends up
negotiating TLS 1.2 cannot use ML-KEM at all, so the opt-out is not recommended
for production.

ML-KEM availability on TLS 1.3 also depends on the Go runtime: Go 1.24 and later
enable the hybrid X25519MLKEM768 key exchange by default, though that default can
be disabled (for example via `GODEBUG`). TLS 1.3 keeps ML-KEM available; it does
not by itself guarantee it on every connection.

## Affected transports

The floor applies to the three production client TLS configurations:

| Transport | Opt-out option |
| --- | --- |
| Main API client (`Connection`) | `ConnectionBuilder.AllowTLS12(bool)` / `allow_tls12` YAML key |
| Token request client (authentication `TransportWrapper`) | `TransportWrapperBuilder.AllowTLS12(bool)` |
| OAuth2 token exchange (`AuthCodeConfig`) | `AuthCodeConfig.AllowTLS12(bool)` |
| JWKS keys client (authentication `HandlerBuilder`) | `HandlerBuilder.KeysAllowTLS12(bool)` |

When you set `ConnectionBuilder.AllowTLS12`, the option is propagated to both the
main API transport and the separate client used to request OpenID tokens, so a
TLS 1.2-only token endpoint is honored without any extra configuration.

The h2c transport is cleartext by design and is unaffected.

## Default behavior

No configuration is required for the secure default. Every production transport
negotiates a minimum of TLS 1.3:

```go
conn, err := sdk.NewConnectionBuilder().
    Tokens(token).
    Build()
// Minimum TLS version is 1.3; post-quantum key exchange remains available.
```

## Lowering the floor to TLS 1.2

### Main API connection

Using the builder:

```go
conn, err := sdk.NewConnectionBuilder().
    Tokens(token).
    AllowTLS12(true). // opt out to TLS 1.2 for legacy compatibility
    Build()
```

Using a configuration file loaded with `ConnectionBuilder.Load`:

```yaml
# connection config
url: https://api.openshift.com
allow_tls12: true
```

```go
conn, err := sdk.NewConnectionBuilder().
    Load("config.yaml").
    Build()
```

### OAuth2 token exchange (authorization code / PKCE flow)

```go
config := authentication.NewAuthCodeConfig().
    Client(clientID).
    AllowTLS12(true)
```

### JWKS keys client (server-side token verification)

```go
handler, err := authentication.NewHandler().
    Logger(logger).
    KeysURL(keysURL).
    KeysAllowTLS12(true). // opt out for the keys-loading HTTP client
    Next(next).
    Build()
```

## When to use the opt-out

Enable TLS 1.2 only when a required endpoint or intermediary genuinely cannot
negotiate TLS 1.3 — for example, a TLS-terminating middlebox, a corporate proxy,
or legacy infrastructure. A failed TLS 1.3 handshake on its own does **not** prove
a version limitation: network errors, SNI/certificate problems, or other
handshake failures can look the same. Before enabling the opt-out, confirm both
that TLS 1.3 fails *specifically* on protocol version and that TLS 1.2 actually
succeeds against the same endpoint:

```bash
# 1. Probe TLS 1.3:
openssl s_client -connect api.openshift.com:443 -tls1_3 </dev/null 2>&1 \
  | grep -iE "Protocol|Cipher|alert|handshake failure"
# Protocol: TLSv1.3        -> endpoint supports 1.3; no opt-out needed.

# 2. Only if step 1 fails, probe TLS 1.2 to confirm the endpoint works there:
openssl s_client -connect api.openshift.com:443 -tls1_2 </dev/null 2>&1 \
  | grep -iE "Protocol|Cipher|alert|handshake failure"
# Protocol: TLSv1.2        -> opt-out is justified.
# also fails / no response -> the problem is NOT a TLS version limitation;
#                             do not enable the opt-out, investigate further.
```

Look for a protocol-version alert (for example `protocol version` / alert 70) on
the TLS 1.3 probe as the signal that the version is the cause. Prefer scoping the
opt-out to the specific connection or client that needs it rather than enabling it
globally, and remove it once the endpoint supports TLS 1.3.

## Security note

Lowering the floor to TLS 1.2 permits connections to negotiate TLS 1.2. Any
connection that does negotiate TLS 1.2 lacks post-quantum (ML-KEM) key exchange
and is exposed to "harvest now, decrypt later" risk; TLS 1.3-capable servers are
unaffected and still negotiate 1.3. Treat the opt-out as a temporary
compatibility measure, not a default.
