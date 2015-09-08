---
permalink: /post-your-guide/
title: Post your guide
---
Work your way through these steps to set up automated publishing to [18F
Pages](https://pages.18f.gov/) for your new guide:

- [Create the `18f-pages` branch.](#create-18f-pages-branch)
- [Set the default branch.](#set-default-branch)
- [Create the publishing webhook.](#set-webhook)
- [Trigger a build.](#trigger-a-build)
- [Add the new guide to 18F Guides.](#add-new-guide)

## <a name="create-18f-pages-branch"></a>Create the `18f-pages` branch

**If you ran the `./go create_repo` command from the _GitHub setup_ chapter,
you can skip ahead to the [Create the publishing webhook](#set-webhook)
section.  Otherwise, keep reading.**

To publish your guide automatically to `pages.18f.gov`, you'll need
to create an `18f-pages` branch. You can do this using the GitHub interface by
clicking the **branch: master** button and entering `18f-pages` in the **Switch
branches/tags** drop-down box:

<img src="{{site.baseurl}}/images/18f-pages.png" alt="GitHub branch creation
interface">

## <a name="set-default-branch"></a>Set the default branch

_Note: If your repository is not just a Jekyll site — for example, if it's a project
repository with a `gh-pages` or `18f-pages` branch for documentation — you can
ignore this step._

You also need to set `18f-pages` branch as the default. First, click the **Settings** page button on the right side of the screen:<br/>
<img src="{{site.baseurl}}/images/gh-settings-button.png" alt="GitHub settings page button">

This will present you with the **Options** page. In the **Settings** section, select `18f-pages` from the **Default branch** drop-down menu:<br/>
<img src="{{site.baseurl}}/images/gh-default-branch.png" alt="GitHub default branch option">

Deleting the original `master` branch, both on GitHub and locally, is left as
an exercise for the reader. Doing so will help avoid confusion in the long run
but isn't strictly necessary.

## <a name="set-webhook"></a>Create the publishing webhook

Go into the **Webhooks & Services** section of the **Settings** section
and click the **Add webhook** button. On the following screen, set the
**Payload URL** to `https://pages.18f.gov/deploy`, leave the **Secret** field
blank, and click **Update webhook**:

<img src="{{site.baseurl}}/images/gh-webhook.png" alt="Set the GitHub webhook">

## <a name="trigger-a-build"></a>Trigger a build

With the webhook in place, push any update to your `18f-pages` branch to your
GitHub repository. Within seconds, your guide should appear at
`https://pages.18f.gov/MY-NEW-GUIDE`. Your guide is now live!

## <a name="add-new-guide"></a>Add the new guide to 18F Guides

You've reached the final step! Add an entry to the `navigation:` list of [18F
Guides](http://18f.github.io/guides/) linking to your new guide. You can [use
this link to edit the file directly in
GitHub](https://github.com/18F/guides/edit/18f-pages/_config.yml):

<img src="{{site.baseurl}}/images/gh-add-guide.png" alt="Add the new guide to 18F Guides using the GitHub text editor">

Congratulations! Your guide should now be published and accessible to the world
as one of the few, the proud, the [18F Guides](https://pages.18f.gov/guides/)!
