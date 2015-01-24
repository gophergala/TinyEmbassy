/*global $*/
define(
  ['marionette','vent','views/header','views/Footer','views/LeftPane','views/RightPane','views/Article'],
  function(marionette, vent, Header, Footer, LeftPane, RightPane, Article){
    "use strict";

    var app = new marionette.Application();

    app.addRegions({
      header : '#header',
      article   : '#article',
      leftpane : '#left-pane',
      rightpane : '#right-pane',
      footer : '#footer'
    });

    app.addInitializer(function(){

        app.header.show(new Header());
        app.leftpane.show(new LeftPane());
        app.article.show(new Article({'title':'hi'}));
        app.rightpane.show(new RightPane());
        app.footer.show(new Footer());
    });
  return app;
  }
);
