/**
 * Created by deepak on 1/25/15.
 */

define(['marionette','templates'],function(Marionette,templates){
    "use strict";

    return Marionette.ItemView.extend({
        template: templates.publishersBadges,
        events:{
            "click .badgeThumb":function(){

            }

        },
        render:function(campaign){
            this.$el.html(this.template({campaign:campaign.attributes}));
        }
    });
});
