## Welcome!

We're so glad you're thinking about contributing to an 18F open source project! If you're unsure or afraid of anything, just ask or submit the issue or pull request anyways. The worst that can happen is that you'll be politely asked to change something. We appreciate any sort of contribution, and don't want a wall of rules to get in the way of that.

Before contributing, we encourage you to read our CONTRIBUTING policy (you are here), our LICENSE, and our README, all of which should be in this repository. If you have any questions, or want to read more about our underlying policies, you can consult the 18F Open Source Policy GitHub repository at https://github.com/18f/open-source-policy, or just shoot us an email/official government letterhead note to [18f@gsa.gov](mailto:18f@gsa.gov).

# Contributing to Compliance Masonry

## Help Wanted On
- Adding new renderers `/src/renderers`
- Developing CLI `/src/cli.py`
- Developing import interface `masonry.py`
- Adding tests
- Documentation


## Setting up development environment
1. Fork Compliance Masonry
2. Clone forked repository
```
git clone https://github.com/<<Your Forked Repo>>.git
```
3. Install the development requirements
```
pip install -r requirements.txt
```
4. Install Compliance Masonry (will need to repeat to see changes)
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
Create a new issue on [our Issues page](https://github.com/opencontrol/compliance-masonry/issues)


## Public domain

This project is in the public domain within the United States, and
copyright and related rights in the work worldwide are waived through
the [CC0 1.0 Universal public domain dedication](https://creativecommons.org/publicdomain/zero/1.0/).

All contributions to this project will be released under the CC0
dedication. By submitting a pull request, you are agreeing to comply
with this waiver of copyright interest.
