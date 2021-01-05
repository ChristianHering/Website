Website
===========

This repository holds a the source for [ChristianHering.com](https://ChristianHering.com/) and it's respective subdomains.

It provides:

  * An outline for a basic webapp written in go

Table of Contents:

  * [About](#about)
  * [Installing and Compiling from Source](#installing-and-compiling-from-source)
  * [Contributing](#contributing)
  * [License](#license)

About
-----

This project is broken up into multiple packages each serving a different purpose. Some links will be broken during the initial development stages

  * The [Admin](/admin/README.md) package serves an administrative panel offering the user statistics and configuration options. Hosted at [https://admin.ChristianHering.com/](https://admin.ChristianHering.com/).
  * The [Blog](/blog/README.md) package serves a simple blog to track my development decisions in a more texual format. Hosted at [https://blog.ChristianHering.com/](https://admin.ChristianHering.com/).
  * The [Consulting](/consulting/README.md) package serves a mostly static site to represent a site your adverage business might have. Hosted at [https://ChristianHering.com/](https://admin.ChristianHering.com/).
  * The [Docs](/docs/README.md) package hosts technical documentation for any API's used by the frontend. (OpenAPI/Swagger) Hosted at [https://docs.ChristianHering.com/](https://admin.ChristianHering.com/).
  * The [Portfolio](/portfolio/README.md) package serves a personal/portfolio website used for showcasing my projects/progress. Hosted at [https://portfolio.ChristianHering.com/](https://admin.ChristianHering.com/).
  * The [Utils](/utils/README.md) package holds utility and initialization functions for the main webapp. This includes configuration/secret management, SQL connection management, middleware, etc

For questions about a specific package, please read their respective README file.

Installing and Compiling from Source
------------

The easiest way to view the current release is to visit [ChristianHering.com](https://ChristianHering.com/). (after v1.0.0 is released)


If you're looking to compile from source, you'll need the following:

  * [Go](https://golang.org) installed and [configured](https://golang.org/doc/install)
  * A [MySQL Database](https://www.mysql.com/) configured with the necessary rows/columns
  * A local install of [NPM](https://www.npmjs.com/) and [VueJS](https://vuejs.org/) for the admin panel
  * Port 80 availible on your testing machine (You can change the port in [main.go](/main.go))
  * A little patience :)

Contributing
------------

Contributions are always welcome. If you're interested in contributing, send me an email or submit a PR.

License
-------

Currently closed source until I choose a license for this project. Please refer to the [license](/LICENSE) file for more information.
