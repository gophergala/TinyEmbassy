<script>
    genCode = function(cmpId,pubId,grpId,bdgId){
        $(".badgeCode").show();


        $(".badgeCode").html(cmpId+"/"+pubId+"/"+grpId+"/"+bdgId);
    }
</script>

<div class="badgeCategories">
    <% _.each(campaign.groups, function(group){ %>
    <div class="badgeContainer">
        <div><%= group.title %></div>
        <% _.each(group.badges, function(badge){ %>
        <div class="badgeImages">
            <img onclick="genCode('<%=campaign.campaign_id %>','<%=campaign.publisher_id%>','<%=group.group_id%>','<%=badge.badge_id%>')"
                 class="badgeThumb"src=".//..//resources//<%= badge.image %>"
                 width=55 />

        </div>
        <% }); %>
    </div>
    <% }); %>
</div>

<div class="badgeCode">
</div>