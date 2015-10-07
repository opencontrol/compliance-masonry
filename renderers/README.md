These scripts converting controls, certification, and standards documentation into different formats. These scripts are intended to be run from the parent directory.

#### workbook_to_yaml.py
Converts `data/xlsx/Control-masonry.xlsx` workbook into yaml files.
```
python renderers/workbook_to_yaml.py
```

#### yamls_to_certification.py
Converts files in `data/controls`, `data/certifications`, `data/standards` into certification documentation and exports files to `/completed_certifications`
```
python renderers/yamls_to_certification.py
```
