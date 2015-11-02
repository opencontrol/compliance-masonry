# Contributing to Control Masonry

## Help Wanted On
- Adding new renderers `/src/renderers`
- Developing CLI `/src/cli.py`
- Developing import interface `masonry.py`
- Adding tests
- Documentation


## Setting up development environment
1. Fork Control Masonry
2. Clone forked repository
```
git clone https://github.com/<<Your Forked Repo>>.git
```
3. Install the development requirements
```
pip install -r requirements.txt
```
4. Install Control Masonry (will need to repeat to see changes)
```
python setup.py install
```
5. Make Changes!

## Running Tests
Before any contributions are accepted changes will need pass existing tests or make changes to the test when needed.
```
tox
```

## Submitting a Pull Request
1. Make sure tests pass
2. Write descriptive commit messages
3. Push changes to forked branch
4. Open pull request to master from forked branch
5. Explain changes in detail


## Found an Issue? or Want a Feature?
Create a new issue on [our Issues page](https://github.com/18F/control-masonry/issues)
