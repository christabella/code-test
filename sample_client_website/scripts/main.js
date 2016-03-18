var timer = null,
    time = 0,
    width = String($(window).width()),
    height = String($(window).height());

if (!Cookies.get('sessionId')) {
    var sessionId = "";
    for (i = 0; i < 21; i++) {
        sessionId += Math.floor(Math.random()*10);
        if ((i == 5) || (i == 12)) {
            sessionId += "-";
        }
    }
    Cookies.set('sessionId', sessionId);
} else {
    var sessionId = Cookies.get('sessionId');
}

/** WINDOW RESIZE EVENT **/

$(window).one("resize", function() {
    $.ajax({
        method: "POST",
        url: "http://localhost:8080",
            data: JSON.stringify({ 
            eventType: "windowResize",
            websiteUrl: window.location.origin,
            sessionId: sessionId,
            widthBefore: width,
            heightBefore: height,
            widthAfter: String($(window).width()),
            heightAfter: String($(window).height()),
        })
    })
    .done(function(msg) {
        console.log("widthAfter: " + $(window).width());
    });

});

/*** COPY/PASTE EVENT ***/

$(".form-control").bind({
    copy : function(event){
        $.ajax({
            method: "POST",
            url: "http://localhost:8080",
            data: JSON.stringify({ 
                eventType: "copyAndPaste",
                websiteUrl: window.location.origin,
                sessionId: sessionId,
                pasted: false, 
                formId: event.target.id
            })
        })
        .done(function(msg) {
            console.log("Copied!: " + event.target.id);
        });

    },
    paste : function(){
        $.ajax({
            method: "POST",
            url: "http://localhost:8080",
            data: JSON.stringify({ 
                eventType: "copyAndPaste",
                websiteUrl: window.location.origin,
                sessionId: sessionId,
                pasted: true, 
                formId: event.target.id
            })
        })
        .done(function(msg) {
            console.log("Pasted!: " + event.target.id);
        });
    }
});

/** TIME TAKEN EVENT **/

$(".form-control").bind("keydown", function(event) {
    start = new Date().getTime();
    console.log(start);
    console.log("Start timing.");
    setTimeout(updateTimer, 100);
    $(".form-control").unbind("keydown");
});

$(".form").submit(function(event) {
    clearTimeout(timer);
    console.log(elapsed);

    $.ajax({
        method: "POST",
        url: "http://localhost:8080",
            data: JSON.stringify({ 
            eventType: "timeTaken",
            websiteUrl: window.location.origin,
            sessionId: sessionId,
            time: elapsed
        })
    })
    .done(function(msg) {
        console.log("TimeTaken: " + elapsed);
    });

});

function updateTimer() 
{  
    time += 100;  
    elapsed = Math.floor(time / 100) / 10;  
    if(Math.round(elapsed) == elapsed) { 
        elapsed += '.0'; 
    }
    // account for possible inaccuracy
    var diff = (new Date().getTime() - start) - time;  
    timer = setTimeout(updateTimer, (100 - diff));  
}  
