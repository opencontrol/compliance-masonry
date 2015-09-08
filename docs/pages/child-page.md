---
permalink: /add-a-new-page/make-a-child-page/
title: Make a child page
parent: Add a new page
---

If you want to nest a page under a parent page, follow the instructions to [add a new page]({{ site.baseurl }}/add-a-new-page/) with two additions to the YAML front matter. Here is the front-matter for this page:

```yaml
---
permalink: {{ page.permalink }}
title: {{ page.title }}
parent: {{ page.parent }}
---
```

Note the `/parent/child/` format for the permalink, and the `parent:`
property. This way, when you're on a parent or child page, the children are visible in the menu. (You'll need to
[run `./go update_nav`]({{ site.baseurl }}/update-the-config-file/#register-new-pages)
before the changes to the menu appear— read more about that in the _Update the
config file_ chapter.)

## Next steps

Click the _Add images_ entry in the table of contents to learn how to add
images to your guide.
