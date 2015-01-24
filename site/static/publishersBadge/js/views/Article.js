/**
 * Created by deepak on 4/6/14.
 */

define(['marionette','templates','./../views/PublisherBadgesView',
'./../models/BadgeCampaign'],
    function(Marionette,templates,PublishersBadgesView,PublisherBadgeModel){
    "use strict";

    return Marionette.ItemView.extend({
        template: templates.article,
        events:{
            "click .emailSubmit":function(){
                var bm  = new PublisherBadgeModel();
                bm.fetch({
                    success: function (campaign) {
                        $('.publisherLogin').hide();
                        console.log(campaign.toJSON());
                        var publishersBadgesView = new PublishersBadgesView();
                        publishersBadgesView.render(campaign);
                        $('#badgeContent').html("");
                        $('#badgeContent').html(publishersBadgesView.$el);
                    }
                });
            }

        }
    });
});
