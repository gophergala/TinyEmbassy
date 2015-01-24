require.config({
  paths : {
    underscore : 'lib/underscore',
    backbone   : 'lib/backbone',
    marionette : 'lib/backbone.marionette',
    jquery     : 'lib/jquery.min',
    jqueryui   : 'lib/jquery-ui',
    tpl        : 'lib/tpl'
  },
  shim : {
      'jquery': {
          exports: '$'
      },

    'lib/backbone-localStorage' : ['backbone'],
    underscore : {
      exports : '_'
    },
    backbone : {
      exports : 'Backbone',
      deps : ['jquery','underscore','jqueryui']
    },
    marionette : {
      exports : 'Backbone.Marionette',
      deps : ['backbone']
    }
  },
  deps : ['jquery','underscore']
});

require(['app','backbone','./models/BadgeCampaign'],function(app,Backbone,BadgeModel){
  "use strict";
  app.start();
  app.on("start",function(options){

});


 // Backbone.history.start();
});
