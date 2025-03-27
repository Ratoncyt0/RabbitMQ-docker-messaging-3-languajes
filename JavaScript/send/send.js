#!/usr/bin/env node

var amqp = require('amqplib/callback_api');

// variables de ambiente :3
const MSG = process.env.MSG || 'error al recibir el mensaje';
const IP = process.env.IP || '172.17.0.1' // ip default de docker

// Envio de mensaje amqp
amqp.connect(`amqp://${IP}`, function(error0, connection) {
    if (error0) {
        throw error0;
    }
    connection.createChannel(function(error1, channel) {
        if (error1) {
            throw error1;
        }

        var queue = 'hello';
        //var msg = 'Hello World!';

        channel.assertQueue(queue, {
            durable: false
        });
        channel.sendToQueue(queue, Buffer.from(MSG));

        console.log(" [x] Sent %s", MSG);
    });
    setTimeout(function() {
        connection.close();
        process.exit(0);
    }, 500);
});

