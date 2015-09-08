var fs = require('fs'),
    yaml = require('js-yaml');

var main_ctrl = yaml.safeLoad(fs.readFileSync(__dirname + '/controls/final_controler.yaml'));
var config_file = yaml.safeLoad(fs.readFileSync(__dirname + '/docs/_config.yml'));


function create_markdown(ctrl_key, title) {
    var markdown = '---\n';
    markdown += 'permalink: /' + ctrl_key + '/\n';
    markdown += 'title: ' + ctrl_key + ' - ' + title + '\n';
    markdown += '---\n';
    return markdown
}

function prepare_justification(justification) {
    if (justification.text) {
        return '* ' + justification.text + '  \n';
    };
};

function build_documentation(control, markdown) {
    Object.keys(control).forEach(function (section_key) {
        if (section_key !== 'name' && control[section_key]) {
            control[section_key].forEach(function (element) {
                markdown += '### ' + element.title + "  \n";
                element.justifications.forEach(function (justification) {
                    markdown += prepare_justification(justification);
                });
                markdown += '  \n';
            });
        };
    })
    return markdown;
};

function write_control(control, ctrl_key) {
    // Create the markdown file
    markdown = create_markdown(ctrl_key, control.name);
    // Add the Documentation
    markdown = build_documentation(control, markdown);
    // Write control to main yaml file
    fs.writeFile(__dirname + '/docs/pages/' + ctrl_key + '.md', markdown, function(err) {
        if (err) {
            return console.log(err);
        };
        console.log("The file was saved!");
    });

};

function update_config(ctrl_key) {
    var page_json = {
        text: ctrl_key,
        url: ctrl_key + '/',
        internal: true
    };
    var page_ref = config_file.navigation.filter(function(page, index) {
        if (page.text === ctrl_key) {
            config_file.navigation[index] = page_json;
            return index;
        };
    });
    if (page_ref.length === 0) {
        config_file.navigation.push(page_json);
    };
};

function render_control(controls) {
    // Loop through base control main keys
    Object.keys(controls).forEach(function(ctrl_key) {
        // Write control to pages repository
        write_control(controls[ctrl_key], ctrl_key);
        // Add page to config file
        update_config(ctrl_key);
    });
    fs.writeFile(__dirname + '/docs/_config.yml', yaml.safeDump(config_file), function(err) {
        if (err) {
            return console.log(err);
        };
        console.log("The config file was saved!");
    });
};


render_control(main_ctrl);
