
Cider
=====

> Status: in initial development

Cider is container based CI/CD tool. The main idea behind it is single responsibility for each container type. There are following types of nodes:

  - **Cider server** - manages the build queue and events and help other services to communicate.
  - **Listeners** - services that discover projects and report new build target to Cider server.
  - **Workers** - services that receive build tasks from Cider server and executes them.
  - **Notifiers** - listen to specific Cider server events and broadcast them to specific channels

## Development team configuration

In a sample configuration Cider cluster could have following roles:
  - Cider server - is a central point of communication.
  - Github listener - authenticates within organization account and discover repositories and build targets, reports them to the server. Subscribes for post-commit hook and after receiving one translates and transmits an event to the server.
  - Cider server - schedules build task.
  - Github worker - receives the build task, checks out the code from Github, executes build task and reports back to the server.
  - Cider server - updates build status and broadcast event to notifiers.
  - Github notifier - sends check status to PR.
  - Slack notifier - publishes message to a slack channel.

## Local configuration

As another example you can run cider locally:
  - Cider server - is a central point of communication.
  - Filesystem listener discovers projects in the configured folder, monitors and reports new commits to the server.
  - Cider server - schedules build task.
  - Local worker - checks out the code to a separate folder, performs the build.
  - Local notifier - in case of build failure triggers system notification.
