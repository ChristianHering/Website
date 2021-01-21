<template>
    <div id="app">
        <div class="row body-grid-ch">
            <div class="col-1"></div>

            <div class="col-xs-12 col-sm-3">
                <ul>
                    <li><router-link to="/">Home</router-link></li>
                    <li><router-link to="/errors">Errors</router-link></li>
                    <li><router-link to="/posts">Blog Posts</router-link></li>
                    <li><router-link to="/postCreate">Create Blog Post</router-link></li>
                    <li><a href="http://admin.ChristianHering.com/logout">Logout</a></li>
                </ul>
            </div>
                <div class="col-xs-12 col-sm-6">
                    <router-view /> <!-- Where selected components are rendered -->
                </div>

            <div class="col-2"></div>
        </div>
    </div>
</template>

<style scoped>
    ul {
        list-style-type: none;
        padding-left: 0px;
    }

    li {
        margin-top: 5px;
        margin-bottom: 5px;
    }
    
    a {
        color: rgb(153, 153, 153);
        text-shadow: 2px 3px 50px rgba(0,0,0,0.5);
        font-size: 18px;
    }
</style>

<style>
    body, html {
        height: 100%;
        background-color: #1f1f1f;
    }

    .fullscreen-ch {
        height: 100vh;
        width: 100%;
    }

    .body-grid-ch {
        color: white;
        text-align: center;
        justify-content: center;
        padding-top: 50px;
        padding-bottom: 50px;
        margin: 0px;
    }

    a {
        text-decoration: none;
    }
</style>
