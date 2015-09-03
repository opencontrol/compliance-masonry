/* This script combines multiple control documentation into one JSON file */
var fs = require('fs'),
    glob = require("glob"),
    yaml = require('js-yaml');

// Load the base control
var base_ctrl = yaml.safeLoad(fs.readFileSync(__dirname + '/controls/base_control.yml'));

// Function for adding a single control to the base control
function add_control(control) {
    // Loop through base control main keys
    Object.keys(base_ctrl).forEach(function(ctrl_key) {
        // Loop through base control sub-keys
        Object.keys(base_ctrl[ctrl_key]).forEach(function(section_key) {
            // Only add control section to base control if it exists
            if (section_key !== 'name' && control[ctrl_key] && control[ctrl_key][section_key]) {
                // create a new array object in base control if it doesn't have one                
                if (!base_ctrl[ctrl_key][section_key]) {
                    base_ctrl[ctrl_key][section_key] = [];
                };
                // Loop through each control section and add to base control
                control[ctrl_key][section_key].forEach(function(control_section) {
                    base_ctrl[ctrl_key][section_key].push(control_section);
                });
            };
        });
    });
};

// 
function create_control() {
    // Find the control files
    glob(__dirname + '/controls/systems/*.yaml', null, function(er, files) {
        files.forEach(function(file) {
            add_control(yaml.safeLoad(fs.readFileSync(file)));
        });
        // Write control to main yaml file
        fs.writeFile(__dirname + '/controls/final_controler.yaml', yaml.safeDump(base_ctrl), function(err) {
            if (err) {
                return console.log(err);
            };
            console.log("The file was saved!");
        });
    });
};

create_control();
