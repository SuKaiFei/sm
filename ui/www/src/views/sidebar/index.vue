<template>
    <div>
        <div class="sidebar-group">
            <b-card
                    no-body
                    img-src="https://placekitten.com/380/200"
                    img-alt="Image"
                    img-top>
                <b-card-body>
                    <b-card-title>小花猫</b-card-title>
                    <b-card-sub-title class="mb-2">精通卖萌，求铲屎官</b-card-sub-title>
                    <b-card-text>
                        头圆、颜面部短，前肢五指，后肢四趾，趾端具锐利而弯曲的爪，爪能伸缩。夜行性。
                    </b-card-text>
                </b-card-body>

                <b-card-body>
                    <a href="#" class="card-link">发布帖子</a>
                    <a href="#" class="card-link">我的帖子</a>
                </b-card-body>
            </b-card>
        </div>
        <div class="sidebar-group box_white">
            <div class="top">
                <h3 class="title"><i class="glyphicon glyphicon-user"></i>&nbsp;<a href="/users" target="_blank"
                                                                                   title="点击更多">新会员</a></h3>
            </div>
            <div class="sb-content user-list" style="clear: both;" data-limit="9">
                <ul style="list-style: none;margin: 0 0 0 0;padding: 0 0 0 0;">
                    <li class="user">
                        <a>
                            <div class="avatar">
                                <img src="https://placekitten.com/g/400/450" alt="Image">
                            </div>
                            <div>
                                <span class="nick-name">小花猫</span>
                            </div>
                        </a>
                    </li>
                    <li class="user">
                        <a>
                            <div class="avatar">
                                <img src="https://placekitten.com/g/400/450" alt="Image">
                            </div>
                            <div>
                                <span class="nick-name">小花猫</span>
                            </div>
                        </a>
                    </li>
                    <li class="user">
                        <a>
                            <div class="avatar">
                                <img src="https://placekitten.com/g/400/450" alt="Image">
                            </div>
                            <div>
                                <span class="nick-name">小花猫</span>
                            </div>
                        </a>
                    </li>
                    <li class="user">
                        <a>
                            <div class="avatar">
                                <img src="https://placekitten.com/g/400/450" alt="Image">
                            </div>
                            <div>
                                <span class="nick-name">小花猫</span>
                            </div>
                        </a>
                    </li>
                    <li class="user">
                        <a>
                            <div class="avatar">
                                <img src="https://placekitten.com/g/400/450" alt="Image">
                            </div>
                            <div>
                                <span class="nick-name">小花猫</span>
                            </div>
                        </a>
                    </li>
                    <li class="user">
                        <a>
                            <div class="avatar">
                                <img src="https://placekitten.com/g/400/450" alt="Image">
                            </div>
                            <div>
                                <span class="nick-name">小花猫</span>
                            </div>
                        </a>
                    </li>
                    <li class="user">
                        <a>
                            <div class="avatar">
                                <img src="https://placekitten.com/g/400/450" alt="Image">
                            </div>
                            <div>
                                <span class="nick-name">小花猫</span>
                            </div>
                        </a>
                    </li>
                    <li class="user">
                        <a>
                            <div class="avatar">
                                <img src="https://placekitten.com/g/400/450" alt="Image">
                            </div>
                            <div>
                                <span class="nick-name">小花猫</span>
                            </div>
                        </a>
                    </li>
                    <li class="user">
                        <a>
                            <div class="avatar">
                                <img src="https://placekitten.com/g/400/450" alt="Image">
                            </div>
                            <div>
                                <span class="nick-name">小花猫</span>
                            </div>
                        </a>
                    </li>
                </ul>
            </div>
        </div>
        <div class="sidebar-group box_white">
            <div class="top">
                <h3 class="title">
                    <i class="glyphicon glyphicon-user"></i>&nbsp;
                    <a href="#" target="_blank" title="点击更多">加入我们</a></h3>
            </div>
            <div class="sb-content">
                QQ群：331497557
            </div>
        </div>
    </div>
</template>

<script>
	export default {
		name: "Sidebar"
	}
</script>

<style type="scss" scoped>
    .sidebar-group {
        margin-bottom: 12px;
    }

    .box_white {
        background: #FFF;
        overflow: hidden;
        border-radius: 0.25rem;
        border: 1px solid rgba(0, 0, 0, 0.125);
    }

    .user-list .user {
        text-align: center;
        margin-bottom: 8px;
        width: 33%;
        float: left;
    }

    .user-list .user .avatar img {
        border-radius: 20%;
        width: 90%;
    }

    .user-list .card-body {
        padding: 0;
    }

    .sidebar-group .top {
        line-height: 24px;
        border-bottom: solid 1px #007bff;
        font-size: 14px;
        font-weight: 700;
        margin-bottom: 4px;
        margin-top: 10px;
        padding-left: 10px;
    }
</style>
