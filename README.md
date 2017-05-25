# Docker Secrets Simulator

![Docker Stars](https://img.shields.io/docker/stars/mikesir87/secrets-simulator.svg)
![Docker Pulls](https://img.shields.io/docker/pulls/mikesir87/secrets-simulator.svg)
![Docker Automated Builds](http://img.shields.io/docker/automated/mikesir87/secrets-simulator.svg)

[Blog post here](https://blog.mikesir87.io/2017/05/using-docker-secrets-during-development/)

Docker Secrets are awesome. They provide a secure way to get secret/sensitive data into your container. However, they require a Swarm to run.  Since secrets are exposed simply as files in `/run/secrets`, we can mock it using one of two methods:

- Mount secrets files to `/run/secrets` from the host (which requires you to define a file per secret)
- Use this image!

The "Secrets Simulator" image will take all defined environment variables and create secrets files where the variable name is the filename and the variable value is the secret content. This allows you to keep everything within your docker-compose.yml!

## How to Use

Here are the quick steps...

1. Add a service using the `mikesir87/secrets-simulator` image to your `docker-compose.yml` file.
2. Define an environment variable for each secret you want to expose in your app.
3. Create a volume that will share the new secrets with your app and mount it to both the simulator and app services.
4. You're done!

## Example

This is the `docker-compose.yml` found in the repo. So... give it a shot yourself!

```yaml
version: '3.1'

services:
  secrets-simulator:
    image: mikesir87/secrets-simulator
    environment:
      DB_USERNAME: admin
      DB_PASSWORD: password1234!
    volumes:
      - secrets:/run/secrets:rw

  viewer:
    image: mikesir87/secrets-viewer
    volumes:
      - secrets:/run/secrets:ro

volumes:
  secrets:
    driver: local
```

After running `docker-compose up`, we have output that looks like this:

```
$ docker-compose up
Creating network "secretsimulator_default" with the default driver
Starting secretsimulator_secrets-simulator_1 ...
Starting secretsimulator_secrets-simulator_1
Starting secretsimulator_viewer_1 ...
Starting secretsimulator_viewer_1 ... done
Attaching to secretsimulator_secrets-simulator_1, secretsimulator_viewer_1
viewer_1             | Starting to dump secrets...
viewer_1             | DB_PASSWORD : password1234!
viewer_1             | DB_USERNAME : admin
secretsimulator_secrets-simulator_1 exited with code 0
secretsimulator_viewer_1 exited with code 0
```

So... what happened?  The contents of `/run/secrets` ended up looking like this...

```
/run/secrets
    DB_PASSWORD      => contents of: password1234!
    DB_USERNAME      => contents of: admin
```



## FAQ

**How can I pick different secrets for different apps?**

You can run as many simulator services as you wish! You'll just need to create a new volume for each distinct set of secrets.


**Why mount secrets as `ro`?**

To help simulate a "real" scenario, I bind the secrets as read-only in the app. That way, I can't make any accidental changes to them.



