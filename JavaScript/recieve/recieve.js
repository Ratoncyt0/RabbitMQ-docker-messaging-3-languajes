#!/usr/bin/env node

var amqp = require('amqplib/callback_api');

// Variables de ambiente
const IP = process.env.IP || '172.17.0.1' // ip de la maquina por default para la version red


amqp.connect(`amqp://${IP}`, function(error0, connection) {
    if (error0) {
        throw error0;
    }
    connection.createChannel(function(error1, channel) {
        if (error1) {
            throw error1;
        }

        var queue = 'hello';

        channel.assertQueue(queue, {
            durable: false
        });

        console.log(" [*] Waiting for messages in %s. To exit press CTRL+C", queue);

        channel.consume(queue, function(msg) {
            console.log(" [x] Received %s", msg.content.toString());
        }, {
            noAck: true
        });
    });
});