const fastify = require('fastify')
const path = require('path')
const assert = require('assert')

async function main() {
    const app = fastify()

    const envUser = process.env.DOC_USER
    const envPassword = process.env.DOC_PASSWORD
    const envPort = process.env.PORT || "3000"

    assert(typeof envUser === 'string', "missing ENV.DOC_USER")
    assert(typeof envPassword === 'string', "missing ENV.DOC_PASSWORD")

    await app.register(require('fastify-basic-auth'), {
        validate(username, password, _req, _reply, done) {
            if (username === envUser && password === password) {
                done()
            } else {
                done(new Error('You can not access'))
            }
        },
        authenticate: true
    })

    app.register(require('fastify-swagger'), {
        routePrefix: '/',
        mode: 'static',
        specification: {
            path: path.join(__dirname, 'specification.yaml')
        },
        uiHooks: {
            onRequest: app.basicAuth
        },
        exposeRoute: true
    })

    app.listen(envPort, '0.0.0.0', (_err, addr) => {
        console.log(addr);
    })
}

main();