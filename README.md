#gotlsinfo
`gotlsinfo` is a simple CLI program that tells you which ciphers are supported by a web server.
It is particularly useful for offline/intranet websites, where you cannot make use of tools like Qualsys SSL Labs[1] or similar services.

`gotlsinfo` currently checks all 317 SSL/TLS ciphers listed by IANA[2].

[1] https://www.ssllabs.com

[2] http://www.iana.org/assignments/tls-parameters/tls-parameters.xml

##Installation
Obtain a copy of the source code and make sure you have a Go compiler installed on your system.

```sh 
git clone https://github.com/jannispinter/gotlsinfo
```

Build the project:

```sh
cd gotlsinfo && go build
```

##Usage
```sh
./gotlsinfo <hostname/ip> <port>
```

###Examples:
```sh
./gotlsinfo google.com 443
./gotlsinfo 127.0.0.1 8443
```

###Only show supported ciphers:
On UNIX, pipe the output of the program to `grep` and exclude "NOT":

```sh
./gotlsinfo google.com 443 | grep -v "NOT"
```

On Windows: Err.. No..

##TODO
- Show the preferred cipher order of the server
- Detect SSL/TLS vulnerabilities like BEAST, CRIME, Heartbleed or POODLE
- Grade Cipher Suite and TLS configuration
- Output advice to strenghten the security of  the TLS configuration

License
------------
`gotlsinfo` is licensed under MIT. There is a copy of the license in the `LICENSE` file.
