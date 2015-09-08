---
permalink: /update-the-config-file/
title: Update the config file
---
Work your way through these steps to update 
the `_config.yml` file — this configures the 18F style template for your specific guide:

- [Set the guide name.](#set-name)
- [Set the `exclude:` entries.](#set-exclude-entries)
- [Register new pages.](#register-new-pages)
- [Update the repository list.](#update-repository-list)
- [Optional: Set the `back_link:` property.](#set-back-link)
- [Optional: Update `google_analytics_ua:`.](#set-google-analytics)

## <a name="set-name"></a>Set the guide name

The `name:` property appears as the guide's overall title. For example:

```yaml
name: {{site.name}}
```

## <a name="set-exclude-entries"></a>Set the `exclude:` entries

Make sure the `exclude:` list contains at least the following files, and add
any other files you might have added that shouldn't appear in the
generated `_site` directory:

```yaml
exclude:
{% for i in site.exclude %}- {{ i }}
{% endfor %}```

## <a name="register-new-pages"></a>Register new pages

The `navigation:` list is used to generate the table of contents. For example,
the `navigation:` section of this guide contains:

```yaml
navigation:
{% for i in site.navigation %}- text: {{ i.text }}
  url: {{ i.url }}
  internal: {{ i.internal }}{% if i.children %}
  children:{% for child in i.children %}
  - text: {{ child.text }}
    url: {{ child.url }}
    internal: {{ child.internal }}{% endfor %}{% endif %}
{% endfor %}```

Run `./go update_nav` from the root directory to update this list
automatically whenever you add pages or make `title:`, `permalink:`, or
`parent:` changes. After running the script, you may want to edit the results
by hand to produce the desired ordering of any new pages; the order of
existing entries will remain the same.

## <a name="update-repository-list"></a>Update the repository list

You'll need to update the `repos:` list to reflect the GitHub
repository that will contain your guide. The first of these repositories
should be the repository for the guide itself; it will be used to generate
the _Edit this page_ and _file an issue_ links in the footer.

The `url:` should be `https://github.com/18F/MY-NEW-GUIDE`, where
`MY-NEW-GUIDE` is the name you gave your clone of the 18F/guides-template
repository. For the `description:` property, it's OK to enter something
generic like "main repository." However, if you aren't certain about either
value, it's also OK to enter placeholder text for these properties and change
them later, ideally before posting to the 18F Pages server. 

The `repos:` entry of this template contains:

```yaml
repos:{% for i in site.repos %}
- name: {{ i.name }}
  description: {{ i.description }}
  url: {{ i.url }}
{% endfor %}
```

## <a name="set-back-link"></a>Optional: set the `back_link:` property

The `back_link:` property produces the _Read more 18F Guides_ link just under
the title of the guide at the top of the page. If your document is not
actually an 18F Guide, you may change this property to link to 18F Pages— or
any other collection of documents to which your new "guide" actually belongs.

## <a name="set-google-analytics"></a>Optional: update `google_analytics_ua:`

The `google_analytics_ua:` property defaults to the Google Analytics account
identifier for all 18F Pages sites. You can override it if you prefer.

## Next steps

Once you're finished updating the config file, click the _GitHub Setup_
entry in the table of contents.
