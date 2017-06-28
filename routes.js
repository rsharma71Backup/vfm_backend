//here only routing is done and if the ro

'use strict';

const register = require('./functions/register');


module.exports = router => {
      
	  router.get('/', (req, res) => res.end('Welcome to p2plending,please hit a service !'));

	 
	router.post('/registerUser', (req, res) => {
        const nid =  Math.floor(Math.random() * (100000 - 1)) + 1;
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
		
		
			
     
		if (!id ||!firstname ||!lastname|| !email ||!password||!operationalemail|| !phone || !relationshipmanageremail ||!customerlimit ||!feepercentage ||!interestearning ||!accountno||!ifsccode ||!pan||!address||
		!id.trim() ||!firstname.trim() ||!lastname.trim()|| !email.trim() ||!password.trim()||!operationalemail.trim()|| !phone.trim() || !relationshipmanageremail.trim() ||!customerlimit.trim() ||!feepercentage.trim() ||!interestearning.trim() ||!accountno.trim()||!ifsccode.trim() ||!pan.trim()||!address.trim())
		 {
             //the if statement checks if any of the above paramenters are null or not..if is the it sends an error report.
			res.status(400).json({message: 'Invalid Request !'});

		} else {
			console.log("register object"+ register)
			
			register.registerUser(id,firstname,lastname,email,password,operational,email,phone,relationshipmanageremail,customerlimit,feepercentage,interestearning,accountno,ifsccode,pan,address)
			.then(result => {

			//	res.setHeader('Location', '/registerUser/'+email);
				res.status(result.status).json({ message: result.message })
			})

			.catch(err => res.status(err.status).json({ message: err.message }));
		}
	});
	router.post('login',function(req,res) {
		console.log(req.body)
    res.send({ "status": "201","usertype": "lender","token": "daidsa876dsa0dslbabds987"})});

}