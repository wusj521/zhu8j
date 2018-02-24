{{define "navbar"}}
<div class="page">
    <div class="page__bd" style="height: 100%;">
        <div class="weui-tab">
            <div class="weui-navbar">
                <a href="/" class="weui-navbar__item weui-bar__item_on">
                    首页
                </a>
                <a href="/sow" class="weui-navbar__item">
                    母猪
                </a>
                <a href="/pigletbirth" class="weui-navbar__item">
                    仔猪
                </a>
                <a href="/pigletout" class="weui-navbar__item">
                    出栏
                </a>
                <a href="/deadpig" class="weui-navbar__item">
                    淘猪
                </a>
                
            </div>
            <div class="weui-tab__panel">

            </div>
        </div>
    </div>
</div>
<script type="text/javascript">
    $(function(){
        $('.weui-navbar__item').on('click', function () {
            $(this).addClass('weui-bar__item_on').siblings('.weui-bar__item_on').removeClass('weui-bar__item_on');
        });
    });
</script>
{{end}}