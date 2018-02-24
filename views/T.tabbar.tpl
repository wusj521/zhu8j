{{define "tabbar"}}

    <div class="page__bd" style="height: 100%;">
        <div class="weui-tab">
            <div class="weui-tab__panel">

            </div>
            <div class="weui-tabbar">
                <a href="/sow" class="weui-tabbar__item weui-bar__item_on">
                    <span style="display: inline-block;position: relative;">
                        <img src="/static/images/icon_tabbar.png" alt="" class="weui-tabbar__icon">
                        <span class="weui-badge" style="position: absolute;top: -2px;right: -13px;">8</span>
                    </span>
                    <p class="weui-tabbar__label">母猪</p>
                </a>
                <a href="/pigletbirth" class="weui-tabbar__item">
                    <img src="/static/images/icon_tabbar.png" alt="" class="weui-tabbar__icon">
                    <p class="weui-tabbar__label">仔猪</p>
                </a>
                <a href="javascript:;" class="weui-tabbar__item">
                    <span style="display: inline-block;position: relative;">
                        <img src="/static/images/icon_tabbar.png" alt="" class="weui-tabbar__icon">
                        <span class="weui-badge weui-badge_dot" style="position: absolute;top: 0;right: -6px;"></span>
                    </span>
                    <p class="weui-tabbar__label">发现</p>
                </a>
                <a href="javascript:;" class="weui-tabbar__item">
                    <img src="/static/images/icon_tabbar.png" alt="" class="weui-tabbar__icon">
                    <p class="weui-tabbar__label">我</p>
                </a>
            </div>            
        </div>
    </div>


<script type="text/javascript">
    $(function(){
        $('.weui-tabbar__item').on('click', function () {
            $(this).addClass('weui-bar__item_on').siblings('.weui-bar__item_on').removeClass('weui-bar__item_on');
        });
    });
</script>
{{end}}