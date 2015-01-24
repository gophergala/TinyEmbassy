/**
 * Created by deepak on 1/24/15.
 */

define(['marionette','templates'],function(Marionette,templates){
    "use strict";

    return Backbone.Model.extend({
        url:"./json/badge-page.json"
    });
});
