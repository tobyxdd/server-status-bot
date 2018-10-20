# Telegram Server Status Bot

[![MIT licensed][1]][2]

[1]: https://img.shields.io/badge/license-MIT-blue.svg
[2]: LICENSE.md

Sample configuration:

    {
        "token": "114514",
        "services": [
            {
                "name": "Shadowsocks",
                "process": "ss-server"
            },
            {
                "name": "MTProxy",
                "process": "mtproto-proxy"
            }
        ]
    }
