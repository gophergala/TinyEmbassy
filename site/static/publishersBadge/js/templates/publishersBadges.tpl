<div class="badgeCategories">
    <% _.each(groups, function(group){ %>
    <div class="badgeContainer">
        <div><%= group.title %></div>
        <% _.each(group.badges, function(badge){ %>
        <div>
            <img class="badgeThumb"src=".//..//resources//<%= badge.image %>" width=55 />
        </div>
        <% }); %>
    </div>
    <% }); %>
</div>