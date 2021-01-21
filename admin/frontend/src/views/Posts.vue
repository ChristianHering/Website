<template>
    <div>

        <h1>Staged Posts</h1>

        <ul>
            <li :key="stagedPost" v-for="stagedPost in this.stagedPosts">
                <button v-on:click="stagedPostUpdate(stagedPost)">Update Post</button>
                <button v-on:click="stagedPostDelete(stagedPost)">Delete Post</button>
                ID: {{stagedPost.ID}} Date: {{stagedPost.Date}} Title: {{stagedPost.Title}}
            </li>
        </ul>

        <div v-if="stagedPosts == null"><h6>None Found.</h6></div>

        <h1>Published Posts</h1>

        <ul>
            <li :key="post" v-for="post in this.posts">
                <a v-bind:href="'https://blog.ChristianHering.com/post/' + post.ID"><button>View Current Post</button></a>
                <button v-on:click="postUpdate(post)">Update Post</button>
                <button v-on:click="postDelete(post)">Delete Post</button>
                ID: {{post.ID}} Date: {{post.Date}} Title: {{post.Title}}
            </li>
        </ul>

        <div v-if="posts == null"><h6>None Found.</h6></div>
    </div>
</template>

<script>
    import axios from "axios"; //For REST Requests

    export default {
        name: "Errors",
        data: function() {
            return {
                stagedPosts: [],
                posts: []
            }
        },
        mounted: function() { //"Main"
            this.refreshStagedPostList()
            this.refreshPostList()
        },
        methods: {
            refreshPostList: function() {
                axios.get("http://admin.christianhering.com/posts")
                    .then((response) => {
                        this.posts = response.data;
                    })
                    .catch(
                        error => console.log(error)
                    );
            },
            postUpdate: function(post) {
                axios.post(
                    "http://admin.christianhering.com/postUpdate",
                    post
                )
                .then((response) => {
                    console.log(response);
                    return true
                }, (error) => {
                    console.log(error);
                    return false
                });

                this.refreshPostList();
            },
            postDelete: function(post) {
                axios.post(
                    "http://admin.christianhering.com/postDelete",
                    post
                )
                .then((response) => {
                    console.log(response);
                    return true
                }, (error) => {
                    console.log(error);
                    return false
                });

                this.refreshPostList();
            },
            refreshStagedPostList: function() {
                axios.get("http://admin.christianhering.com/stagedPosts")
                    .then((response) => {
                        this.stagedPosts = response.data;
                    })
                    .catch(
                        error => console.log(error)
                    );
            },
            stagedPostUpdate: function(stagedPost) {
                axios.post(
                    "http://admin.christianhering.com/stagedPostUpdate",
                    stagedPost
                )
                .then((response) => {
                    console.log(response);
                    return true
                }, (error) => {
                    console.log(error);
                    return false
                });

                this.refreshStagedPostList();
            },
            stagedPostDelete: function(stagedPost) {
                axios.post(
                    "http://admin.christianhering.com/stagedPostDelete",
                    stagedPost
                )
                .then((response) => {
                    console.log(response);
                    return true
                }, (error) => {
                    console.log(error);
                    return false
                });

                this.refreshStagedPostList();
            }
        }
    };
</script>