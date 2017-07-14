//here only routing is done and if the ro

'use strict';

const register = require('./functions/register');
const createProgram = require('./functions/createProgram');


module.exports = router => {

    router.get('/', (req, res) => res.end('Welcome to p2plending,please hit a service !'));


    router.post('/registerUser', (req, res) => {

        const nid = Math.floor(Math.random() * (100000 - 1)) + 1;
        const id = nid.toString();
        const firstname = req.body.firstname;
        const lastname = req.body.lastname;
        const email = req.body.email;
        const password = req.body.password;
        const operationalemail = req.body.operationalemail;
        const phone = req.body.phone;
        const relationshipmanageremail = req.body.relationshipmanageremail;
        const customerlimit = req.body.customerlimit;
        const feepercentage = req.body.feepercentage;
        const interestearning = req.body.interestearning;
        const accountno = req.body.accountno;
        const ifsccode = req.body.ifsccode;
        const pan = req.body.pan;
        const address = req.body.address;

        if (!id || !firstname || !lastname || !email || !password || !operationalemail || !phone || !relationshipmanageremail || !customerlimit || !feepercentage || !interestearning || !accountno || !ifsccode || !pan || !address ||
            !id.trim() || !firstname.trim() || !lastname.trim() || !email.trim() || !password.trim() || !operationalemail.trim() || !phone.trim() || !relationshipmanageremail.trim() || !customerlimit.trim() || !feepercentage.trim() || !interestearning.trim() || !accountno.trim() || !ifsccode.trim() || !pan.trim() || !address.trim()) {
            //the if statement checks if any of the above paramenters are null or not..if is the it sends an error report.
            res.status(400).json({ message: 'Invalid Request !' });

        } else {
          

            register.registerUser(id, firstname, lastname, email, password, operationalemail, phone, relationshipmanageremail, customerlimit, feepercentage, interestearning, accountno, ifsccode, pan, address)
                .then(result => {

                    //	res.setHeader('Location', '/registerUser/'+email);
                    res.status(result.status).json({ message: result.message })
                })

            .catch(function(err) {
                if (err.status >= 100 && err.status < 600) {
                    res.status(err.status).json({ message: err.message });
                } else {
                    res.status(500);
                }

            });
        }
    });

    router.post('/login', function(req, res) {
        console.log(req.body)
        res.send({ "status": "201", "usertype": "lender", "token": "daidsa876dsa0dslbabds987" })
    });
    
     router.post('/createProgram', function(req, res) {
         const nid = Math.floor(Math.random() * (100000 - 1)) + 1;
        const id = nid.toString();
           const manufacturer = req.body.manufacturer;
            const supplier = req.body.supplier;

            if(!manufacturer||!supplier){
             res.status(400).json({ message: 'Invalid Request !' });
            }
            else{
                createProgram.createProgram(id,manufacturer,supplier)
                .then(result => {

                    //	res.setHeader('Location', '/registerUser/'+email);
                    res.status(result.status).json({ message: result.message })
                })

            .catch(function(err) {
                if (err.status >= 100 && err.status < 600) {
                    res.status(err.status).json({ message: err.message });
                } else {
                    res.status(500);
                }

            });
        }
         });
}