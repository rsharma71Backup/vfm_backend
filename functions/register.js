'use strict';

const bc_client = require('../blockchain_sample_client');
const bcrypt = require('bcryptjs');
var bcSdk = require('../src/blockchain/blockchain_sdk');
var user = 'risabh.s';
var affiliation = 'fundraiser';
//exports is used here so that registerUser can be exposed for router and blockchainSdk file
exports.registerUser = (id, firstname, lastname, email, password, operationalemail, phone, relationshipmanageremail, customerlimit, feepercentage, interestearning, accountno, ifsccode, pan, address) =>
    new Promise((resolve, reject) => {


        const newUser = ({
            id: id,
            firstname: firstname,
            lastname: lastname,
            email: email,
            password: password,
            operationalemail: operationalemail,
            phone: phone,
            relationshipmanageremail: relationshipmanageremail,
            customerlimit: customerlimit,
            feepercentage: freepercentage,
            interestearning: interestearning,
            accountno: accountno,
            ifsccode: ifsccode,
            pan: pan,
            address: address
        });

        console.log("ENTERING THE Userregisteration from register.js to blockchainSdk");

        bcSdk.UserRegisteration({ user: user, UserDetails: newUser })
            .then(() => resolve({ status: 201, message: usertype }))
            .catch(err => {

                if (err.code == 11000) {

                    reject({ status: 409, message: 'User Already Registered !' });

                } else {
                    conslole.log("error occurred" + err);

                    reject({ status: 500, message: 'Internal Server Error !' });
                }
            });
    });