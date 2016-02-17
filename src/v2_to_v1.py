UNCHANGED_FIELDS = ['name', 'documentation_complete', 'references']


def add_if_exists(new_data, old_data, field):
    """ Adds the field to the new data if it exists in the old data """
    if field in old_data:
        new_data[field] = old_data.get(field)


def transport_usable_data(new_data, old_data):
    """ Adds the data structures that haven't changed to the new dictionary """
    for field in UNCHANGED_FIELDS:
        add_if_exists(new_data=new_data, old_data=old_data, field=field)


def unflatten_verifications(old_verifications):
    """ Convert verifications from v2 to v1 """
    new_verifications = {}
    for verification in old_verifications:
        key = verification['key']
        del verification['key']
        new_verifications[key] = verification
    return new_verifications


def transform_covered_by(covered_by):
    """ Transform covered_by to references as in version 1.0 """
    new_references = []
    for ref in covered_by:
        new_ref = {}
        new_ref['verification'] = ref['verification_key']
        component_key = ref.get('component_key')
        system_key = ref.get('system_key')
        if component_key:
            new_ref['component'] = component_key
        if system_key:
            new_ref['system'] = system_key
        new_references.append(new_ref)
    return new_references


def unflatten_satisfies(old_satisfies):
    """ Convert satisfies from v2 to v1 """
    new_satisfies = {}
    for element in old_satisfies:
        new_element = {}
        # Handle exsiting data
        add_if_exists(
            new_data=new_element,
            old_data=element,
            field='narrative'
        )
        add_if_exists(
            new_data=new_element,
            old_data=element,
            field='implementation_status'
        )
        # Handle covered_by
        references = transform_covered_by(element.get('covered_by', {}))
        control_key = element['control_key']
        standard_key = element['standard_key']
        if references:
            new_element['references'] = references
        # Unflatten
        if standard_key not in new_satisfies:
            new_satisfies[standard_key] = {}
        if control_key not in new_satisfies[standard_key]:
            new_satisfies[standard_key][control_key] = new_element
    return new_satisfies


def convert(old_data):
    """ Convert the data from the v2 to v1 """
    new_data = {}
    transport_usable_data(new_data=new_data, old_data=old_data)
    verifications = unflatten_verifications(old_data.get('verifications', {}))
    if verifications:
        new_data['verifications'] = verifications
    satisfies = unflatten_satisfies(old_data.get('satisfies', {}))
    if satisfies:
        new_data['satisfies'] = satisfies
    return new_data
