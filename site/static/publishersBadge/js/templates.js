
define(function(require){
  "use strict";
  return {
        header              : require('tpl!templates/header.tmpl'),
        footer              : require('tpl!templates/footer.tmpl'),
        leftpane           :require('tpl!templates/left-pane.tmpl'),
        rightpane           :require('tpl!templates/right-pane.tmpl'),
        article           :require('tpl!templates/article.tmpl'),
      publishersBadges      :require('tpl!templates/publishersBadges.tpl')
  };
});

