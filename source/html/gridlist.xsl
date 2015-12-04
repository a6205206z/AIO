<?xml version="1.0" encoding="utf-8"?>
<xsl:stylesheet version="1.0"
xmlns:xsl="http://www.w3.org/1999/XSL/Transform">

  <xsl:template match="/">
    <!-- paulirish.com/2008/conditional-stylesheets-vs-css-hacks-answer-neither/ -->
    <!--[if lt IE 7]> <html class="no-js ie6 oldie" lang="en"> <![endif]-->
    <!--[if IE 7]>    <html class="no-js ie7 oldie" lang="en"> <![endif]-->
    <!--[if IE 8]>    <html class="no-js ie8 oldie" lang="en"> <![endif]-->
    <!--[if gt IE 8]><!-->
    <html class="no-js" lang="en">
      <head>
        <title>Welcome to AIO ESB!</title>

        <meta charset="utf-8" />
        <!-- Always force latest IE rendering engine (even in intranet) & Chrome Frame -->
        <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1" />
        <!-- Mobile viewport optimized: h5bp.com/viewport -->
        <meta name="viewport" content="width=device-width" />

        <meta name="robots" content="noindex, nofollow" />
        <meta name="description" content="MetroUI-Web : Simple and complete web UI framework to create web apps with Windows 8 Metro user interface." />
        <meta name="keywords" content="metro, metroui, metro-ui, metro ui, windows 8, metro style, bootstrap, framework, web framework, css, html" />
        <meta name="author" content="AozoraLabs by Marcello Palmitessa"/>

        <!-- remove or comment this line if you want to use the local fonts -->
        <link href='http://fonts.googleapis.com/css?family=Open+Sans:300,400,600,700' rel='stylesheet' type='text/css'/>

        <link rel="stylesheet" type="text/css" href="content/css/bootstrap.css"/>
        <link rel="stylesheet" type="text/css" href="content/css/bootstrap-responsive.css"/>
        <link rel="stylesheet" type="text/css" href="content/css/bootmetro.css"/>
        <link rel="stylesheet" type="text/css" href="content/css/bootmetro-tiles.css"/>
        <link rel="stylesheet" type="text/css" href="content/css/bootmetro-charms.css"/>
        <link rel="stylesheet" type="text/css" href="content/css/metro-ui-dark.css"/>
        <link rel="stylesheet" type="text/css" href="content/css/icomoon.css"/>

        <!--  these two css are to use only for documentation -->
        <link rel="stylesheet" type="text/css" href="content/css/demo.css"/>
        <link rel="stylesheet" type="text/css" href="scripts/google-code-prettify/prettify.css" />

        <!-- Le fav and touch icons -->
        <link rel="shortcut icon" href="content/ico/favicon.ico"/>
        <link rel="apple-touch-icon-precomposed" sizes="144x144" href="content/ico/apple-touch-icon-144-precomposed.png"/>
        <link rel="apple-touch-icon-precomposed" sizes="114x114" href="content/ico/apple-touch-icon-114-precomposed.png"/>
        <link rel="apple-touch-icon-precomposed" sizes="72x72" href="content/ico/apple-touch-icon-72-precomposed.png"/>
        <link rel="apple-touch-icon-precomposed" href="content/ico/apple-touch-icon-57-precomposed.png"/>

        <!-- All JavaScript at the bottom, except for Modernizr and Respond.
      Modernizr enables HTML5 elements & feature detects; Respond is a polyfill for min/max-width CSS3 Media Queries
      For optimal performance, use a custom Modernizr build: www.modernizr.com/download/ -->
        <script src="scripts/modernizr-2.6.1.min.js"></script>
        <script type="text/javascript">
          <![CDATA[   
        
      var _gaq = _gaq || [];
      _gaq.push(['_setAccount', 'UA-3182578-6']);
      _gaq.push(['_trackPageview']);
      (function() {
         var ga = document.createElement('script'); ga.type = 'text/javascript'; ga.async = true;
         ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';
         var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(ga, s);
      })();
   
      ]]>
        </script>
      </head>
      <body data-accent="blue">
        <header id="nav-bar" class="container-fluid">
          <div class="row-fluid">
            <div class="span8">
              <div id="header-container">
                <a id="backbutton" class="win-backbutton" href="#"></a>
                <h5>The New ESB</h5>
                <div class="dropdown">
                  <a class="header-dropdown dropdown-toggle accent-color" data-toggle="dropdown" href="#">
                    Start
                    <b class="caret"></b>
                  </a>
                  <ul class="dropdown-menu">
                    <li>
                      <a href="#">Esb Platform</a>
                    </li>
                    <li>
                      <a href="#">ErrorTrack</a>
                    </li>
                    <li class="divider"></li>
                    <li>
                      <a href="./index.xml">Home</a>
                    </li>
                  </ul>
                </div>
              </div>
            </div>
            <div id="top-info" class="pull-right">
              <a href="#" class="pull-left">
                <div class="top-info-block">
                  <h3>FirstName</h3>
                  <h4>LastName</h4>
                </div>
                <div class="top-info-block">
                  <b class="icon-user"></b>
                </div>
              </a>
              <hr class="separator pull-left"/>
              <a id="settings" class="pull-left" href="#">
                <b class="icon-settings"></b>
              </a>
            </div>
          </div>
        </header>
        <div class="container-fluid">
          <div class="row-fluid">
            <div class="metro span12">
              <xsl:apply-templates select="services"/>
            </div>
          </div>
        </div>
        <div id="charms" class="win-ui-dark">
          <div id="theme-charms-section" class="charms-section">
            <div class="charms-header">
              <a href="#" class="close-charms win-command">
                <span class="win-commandimage win-commandring">
                  &#xe05d;
                </span>
              </a>
              <h2>
                Settings
              </h2>
            </div>
            <div class="row-fluid">
              <div class="span12">
                <form class="">
                  <label for="win-theme-select">
                    Change theme:
                  </label>
                  <select id="win-theme-select" class="">
                    <option value="metro-ui-semilight">Semi-Light</option>
                    <option value="metro-ui-light">Light</option>
                    <option value="metro-ui-dark">Dark</option>
                  </select>
                </form>
              </div>
            </div>
          </div>
        </div>

        <script src="//ajax.googleapis.com/ajax/libs/jquery/1.8.2/jquery.min.js"></script>
        <script>
          <![CDATA[
          window.jQuery || document.write("<script src='scripts/jquery-1.8.2.min.js'>\x3C/script>")
          ]]>
        </script>
        <script type="text/javascript" src="scripts/google-code-prettify/prettify.js"></script>
        <script type="text/javascript" src="scripts/jquery.mousewheel.js"></script>
        <script type="text/javascript" src="scripts/jquery.scrollTo.js"></script>
        <script type="text/javascript" src="scripts/bootstrap.min.js"></script>
        <script type="text/javascript" src="scripts/bootmetro.js"></script>
        <script type="text/javascript" src="scripts/bootmetro-charms.js"></script>
        <script type="text/javascript" src="scripts/demo.js"></script>
        <script type="text/javascript" src="scripts/holder.js"></script>
        <script type="text/javascript">
          $(".metro").metro();
        </script>
      </body>
    </html>
  </xsl:template>

  <xsl:template match="services">

    <div class="metro-sections">
      <h2>
        The New ESB Services List
      </h2>
      <div id="section{position()}" class="metro-section tile-span-4">
        <h2>
          Order Service
        </h2>
        <xsl:for-each select="service[contains(name,'Order') or contains(name,'order')]">
          <a target="_blank"  href="{url}">
            <xsl:if test="position() mod 4 = 1 ">
              <xsl:attribute name="class">
                tile wide app wideimage bg-color-orange
              </xsl:attribute>
            </xsl:if>
            <xsl:if test="position() mod 4 = 2 ">
              <xsl:attribute name="class">
                tile app bg-color-blueDark
              </xsl:attribute>
            </xsl:if>
            <xsl:if test="position() mod 4 = 3 ">
              <xsl:attribute name="class">
                tile app bg-color-red
              </xsl:attribute>
            </xsl:if>
            <xsl:if test="position() mod 4 = 0 ">
              <xsl:attribute name="class">
                tile app bg-color-yellow
              </xsl:attribute>
            </xsl:if>

            <xsl:if test="position() mod 5 = 1 ">
              <div class="image-wrapper">
                <img src="content/img/windows_8_logo.png" />
              </div>
            </xsl:if>
            <xsl:if test="position() mod 5 = 2 ">
              <div class="image-wrapper">
                <img src="content/img/wcl.jpg" />
              </div>
            </xsl:if>
            <xsl:if test="position() mod 5 = 3 ">
              <div class="image-wrapper">
                <img src="content/img/tile-wide-collection-1.jpg" />
              </div>
            </xsl:if>
            <xsl:if test="position() mod 5 = 4 ">
              <div class="image-wrapper">
                <img src="content/img/bs-docs-responsive-illustrations.png" />
              </div>
            </xsl:if>
            <div class="column-text">
              <div class="text4">
              </div>
            </div>
            <div class="app-label">
              click here to link <xsl:value-of select="name"/>
            </div>
          </a>
        </xsl:for-each>
      </div>
      <div id="section{position()}" class="metro-section tile-span-4">
        <h2>
          Customer Service
        </h2>
        <xsl:for-each select="service[contains(name,'Cust') or contains(name,'cust')]">
          <a target="_blank"  href="{url}">
            <xsl:if test="position() mod 4 = 1 ">
              <xsl:attribute name="class">
                tile wide app wideimage bg-color-orange
              </xsl:attribute>
            </xsl:if>
            <xsl:if test="position() mod 4 = 2 ">
              <xsl:attribute name="class">
                tile app bg-color-blueDark
              </xsl:attribute>
            </xsl:if>
            <xsl:if test="position() mod 4 = 3 ">
              <xsl:attribute name="class">
                tile app bg-color-red
              </xsl:attribute>
            </xsl:if>
            <xsl:if test="position() mod 4 = 0 ">
              <xsl:attribute name="class">
                tile app bg-color-yellow
              </xsl:attribute>
            </xsl:if>

            <xsl:if test="position() mod 5 = 1 ">
              <div class="image-wrapper">
                <img src="content/img/windows_8_logo.png" />
              </div>
            </xsl:if>
            <xsl:if test="position() mod 5 = 2 ">
              <div class="image-wrapper">
                <img src="content/img/sample1.png" />
              </div>
            </xsl:if>
            <xsl:if test="position() mod 5 = 3 ">
              <div class="image-wrapper">
                <img src="content/img/tile-wide-collection-1.jpg" />
              </div>
            </xsl:if>
            <xsl:if test="position() mod 5 = 4 ">
              <div class="image-wrapper">
                <img src="content/img/bs-docs-responsive-illustrations.png" />
              </div>
            </xsl:if>
            <div class="column-text">
              <div class="text4">
               
              </div>
            </div>
            <div class="app-label">
              click here to link <xsl:value-of select="name"/>
            </div>
          </a>
        </xsl:for-each>
      </div>
      <div id="section{position()}" class="metro-section tile-span-4">
        <h2>
          Other Service
        </h2>
        <xsl:for-each select="service[not(contains(name,'Cust')) and not(contains(name,'cust')) and not(contains(name,'Order')) and not(contains(name,'order'))]">
          <a target="_blank" href="{url}">
            <xsl:if test="position() mod 4 = 1 ">
              <xsl:attribute name="class">
                tile wide app wideimage bg-color-orange
              </xsl:attribute>
            </xsl:if>
            <xsl:if test="position() mod 4 = 2 ">
              <xsl:attribute name="class">
                tile app bg-color-blueDark
              </xsl:attribute>
            </xsl:if>
            <xsl:if test="position() mod 4 = 3 ">
              <xsl:attribute name="class">
                tile app bg-color-red
              </xsl:attribute>
            </xsl:if>
            <xsl:if test="position() mod 4 = 0 ">
              <xsl:attribute name="class">
                tile app bg-color-yellow
              </xsl:attribute>
            </xsl:if>
            
            <xsl:if test="position() mod 5 = 1 ">
              <div class="image-wrapper">
                <img src="content/img/windows_8_logo.png" />
              </div>
            </xsl:if>
            <xsl:if test="position() mod 5 = 2 ">
              <div class="image-wrapper">
                <img src="content/img/sample1.png" />
              </div>
            </xsl:if>
            <xsl:if test="position() mod 5 = 3 ">
              <div class="image-wrapper">
                <img src="content/img/tile-wide-collection-1.jpg" />
              </div>
            </xsl:if>
            <xsl:if test="position() mod 5 = 4 ">
              <div class="image-wrapper">
                <img src="content/img/bs-docs-responsive-illustrations.png" />
              </div>
            </xsl:if>
            
            <div class="column-text">
              <div class="text4">
                
              </div>
            </div>
            <div class="app-label">
              click here to link <xsl:value-of select="name"/>
            </div>
          </a>
        </xsl:for-each>
      </div>
    </div>

  </xsl:template>
</xsl:stylesheet>
