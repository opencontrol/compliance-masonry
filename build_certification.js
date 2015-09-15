/* This script combines controls and standards yamls under a certifications yamls */

// Loading Requirements
var fs = require('fs'),
    glob = require("glob"),
    yaml = require('js-yaml');

// Loop Through individual files and load them into standards, controls, and
// certifications objs
function GetData(callback) {
    var data = {};
    glob(__dirname + '/data/*/*.yaml', null, function(er, files) {
        data.standards = {};
        data.controls = [];
        data.certifications = {}
        files.forEach(function(file) {
            if (file.indexOf('/standards/') !== -1) {
                var temp = yaml.safeLoad(fs.readFileSync(file));
                data.standards[temp.name] = temp;
            } else if (file.indexOf('/controls/') !== -1) {
                data.controls.push.apply(data.controls, yaml.safeLoad(fs.readFileSync(file))['controls']);
            } else if (file.indexOf('/certifications/') !== -1) {
                var temp = yaml.safeLoad(fs.readFileSync(file));
                data.certifications[temp.name] = temp;
            };
        });
        callback(data)
    });
};

GetData(function(data) {
    // Find the standard elements that each control satisfies
    // and append them to standards
    data.controls.forEach(function(control) {
        var satisfies = control['satisfies'];
        Object.keys(satisfies).forEach(function(standard) {
            satisfies[standard].forEach(function(element) {
                if (!data.standards[standard][element].justifications) {
                    data.standards[standard][element].justifications = []
                }
                data.standards[standard][element].justifications.push(control);
            });
        });
    });
    // Create the certification yamls files by merging standards data
    // Loop through certifications
    Object.keys(data.certifications).forEach(function(certification) {
        // Loop through certification standards
        Object.keys(data.certifications[certification].standards).forEach(function(standard) {
            // Loop through each element of the certification standards
            Object.keys(data.certifications[certification].standards[standard]).forEach(function(element) {
                // If the standard element exists in the file append to certifications
                if (data.standards[standard][element]) {
                    data.certifications[certification].standards[standard][element] = (JSON.parse(JSON.stringify(data.standards[standard][element])));;
                }
            });
        });
        // Write certification to main yaml file
        fs.writeFile(__dirname + '/completed_certifications/' + certification + '.yaml', yaml.safeDump(data.certifications[certification]), function(err) {
            if (err) {
                return console.log(err);
            };
            console.log("The file was saved!");
        });

    });

});
