var baseUtils = function(Backbone){

    $(window).scroll(function() {

        var scrollTop = $(window).scrollTop();

        if (scrollTop > 48) {

           $('#header-strip').css({
               'position':'fixed',
               'top':'0%'

           });

            $('#header-strip').hover(function(){
                $('#header-strip').css({
                   


                });
            });
        } else {
            if (scrollTop < 48) {
                $('#header-strip').css({
                    'position':'static '
                });
            }

        }

    });
}