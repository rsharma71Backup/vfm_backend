var cfenv = require('cfenv');
var appEnv = cfenv.getAppEnv();

console.log("server is running on" + appEnv.port)