const { verify } = require("crypto");
const net = require("net");
const { Stream } = require("stream");
const fs = require("fs");

async function main() {
    const client = new OcsaClient({
        host: "localhost",
        port: 8052,
        verbose: true,
    });

    await client.storeFile({
        srcFilePath: "/home/parth/Parth Degama.pdf",
        destFilePath: "hello/p.pdf",
    });
}

class OcsaClient {
    constructor(config) {
        this.config = {
            host: "localhost",
            port: 8052,
            authToken: "qwas",
            tls: false,
            verbose: false,
            ...config
        };
    }

    async storeFile(file) {
        const ocsa = this

        let client = new net.Socket()
        client.connect(ocsa.config.port, ocsa.config.host, function () {
            if (ocsa.config.verbose) {
                console.log('Connected');
            }
        });

        client.on('data', function (data) {
            if (ocsa.config.verbose) {
                console.log('Received: ' + data);
            }

            let d = data.toString()


            switch (d) {
                case "<<<<START_HEADER>>>>":
                    client.write(`FILE_PATH:${file.destFilePath}\nAUTH_TOKEN:${ocsa.config.authToken}\n<<<<END_HEADER>>>>\n`)
                    //client.write(``)
                    //client.write(`<<<<END_HEADER>>>>\n`)
                    break
                case "<<<<START_FILE>>>>":
                    let file_ocsa = fs.createReadStream(file.srcFilePath, { highWaterMark: 1024 * 1024 })

                    file_ocsa.on('data', (chunk) => {
                        client.write(chunk)
                    })

                    file_ocsa.on('end', () => {
                        setTimeout(() => {
                            client.write(`<<<<END_FILE>>>>`)
                        }, 500)
                    })

                    break
                default:
                    throw new Error("Invalid response from server")
                //client.destroy();
            };

            client.on('close', () => {
                if (ocsa.config.verbose) {
                    console.log('file send');
                    process.exit(0)
                }
            })

        })
    }
}
main();
