'use strict';

var bcSdk = require('../src/blockchain/blockchain_sdk');
var user = 'risabh.s';
var affiliation = 'fundraiser';
//exports is used here so that registerUser can be exposed for router and blockchainSdk file
exports.createProgram = (id,manufacturer,supplier) =>
    new Promise((resolve, reject) => {


        const newUser = ({
            id:id,
            manufacturer:manufacturer,
            supplier:supplier
        });

        console.log("ENTERING THE Userregisteration from register.js to blockchainSdk");

        bcSdk.createProgram({ user: user, createProgram: newUser })
            .then(() => resolve({ status: 201, message: "saved program" }))
            .catch(err => {
                        
                    console.log("error occurred" + err);

                    reject({ status: 500, message: 'Internal Server Error !' });
                }
            );
    });