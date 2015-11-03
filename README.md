
A blueprint of plugin implementation in cloud foundry.

Steps:

- Clone cf cli to your workspace `https://github.com/cloudfoundry/cli`
- Compile and generate the executable file from the test.go ``go install``
- Install the plugin by ``cf install-plugin [PATH_TO_PLUGIN_EXECUTABLE]`` command 
- See all plugins with ``cf plugins``


