var express = require('express');
var router = express.Router();
var bodyParser = require('body-parser');
var multer = require("multer");
var upload = multer();
//var type = upload.single('recfile');

var storage = multer.diskStorage({
    destination: (req, file, cb) => {
      cb(null, 'upload')
    },
    filename: (req, file, cb) => {
      cb(null, file.fieldname + '-' + Date.now())
    }
});
var upload = multer({storage: storage});
//var storage = multer.diskStorage({
//    destination: (req, file, cb) => {
 //     cb(null, 'public/images/uploads')
 //   },
 //   filename: (req, file, cb) => {
 //     cb(null, file.fieldname + '-' + Date.now())
 //   }
//});
//var upload = multer({storage: storage});

router.use(bodyParser.urlencoded({ extended: true }));
router.use(bodyParser.json());

var TFBC = require("./FabricHelper")


// Request LC
router.post('/requestLC', function (req, res) {

TFBC.requestLC(req, res);

});

// Issue LC
router.post('/issueLC', function (req, res) {

    TFBC.issueLC(req, res);
    
});

// Accept LC
router.post('/acceptLC', function (req, res) {

    TFBC.acceptLC(req, res);
    
});

// Get LC
router.post('/getLC', function (req, res) {

    TFBC.getLC(req, res);
    
});

// Get LC history
router.post('/getLCHistory', function (req, res) {

    TFBC.getLCHistory(req, res);
    
});

// exporter
//router.post('/exporter', function (req, res) {

//    TFBC.exporter(req, res);
    
//});

router.post('/exporter',function (req, res) {
  //  const file = req.file
   // if (!file) {
   //   const error = new Error('Please upload a file')
    //  error.httpStatusCode = 400
    //  return next(error)
   // }
     TFBC.exporter(req,res);
    
  });


module.exports = router;
