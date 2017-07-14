'use strict';

//const user = require('../blockchai');
var user= "risabh.s";
var bcSdk = require('../src/blockchain/blockchain_sdk.js');


exports.loginUser = (email1, passpin) => 

	new Promise((resolve,reject) => {
		const ui_login =({

			email: email1,
			passpin: passpin

			
		});
		console.log("ENTERING THE login MODULE");
            return bcSdk.User_login({ ui_login})

			.then(() => resolve({ status: 201, message: 'User signed in Sucessfully !',token:token }))

			.catch(err => {

			if (err.code == 11000) {
						
				reject({ status: 409, message: 'User Already Registered !' });

			} else {
				conslole.log("error occurred" + err);

	}
					})
	});

	
	
