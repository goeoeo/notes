#### 静态资源管理
1.静态资源包定义
```php
<?php

namespace app\assets;

use yii\web\AssetBundle;

/**
 * Main application asset bundle.
 *
 * @author Qiang Xue <qiang.xue@gmail.com>
 * @since 2.0
 */
class AppAsset extends AssetBundle
{
    public $basePath = '@webroot';
    public $baseUrl = '@web';
    public $css = [
        'css/site.css',
        'css/main.css',
        'css/red.css',
    ];
    public $js = [
        'js/jquery-migrate-1.2.1.js',
        'js/bootstrap.min.js',
        ['js/html5shiv.js','condition'=>'lte IE9','position'=>\yii\web\View::POS_HEAD],
        ['js/respond.min.js','condition'=>'lte IE9','position'=>\yii\web\View::POS_HEAD],
    ];
    public $depends = [
        'yii\web\YiiAsset',
        'yii\bootstrap\BootstrapAsset',
    ];
}

```
2.使用静态资源包
```php
<?php

/* @var $this \yii\web\View */
/* @var $content string */

use app\widgets\Alert;
use yii\helpers\Html;
use yii\bootstrap\Nav;
use yii\bootstrap\NavBar;
use yii\widgets\Breadcrumbs;
use app\assets\AppAsset;

AppAsset::register($this);
?>
<?php $this->beginPage() ?>
<!DOCTYPE html>
<html lang="<?= Yii::$app->language ?>">
<head>
    <meta charset="<?= Yii::$app->charset ?>">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <?= Html::csrfMetaTags() ?>
    <title><?= Html::encode($this->title) ?></title>
    <?php $this->head() ?>
</head>
<body>
<?php $this->beginBody() ?>
<div class="wrapper">
    <?
        NavBar::begin(['options'=>[
            'class'=>'top-bar animate-dropdown'
        ]]);

        echo Nav::widget([
            'options'=>['class'=>'navbar-nav navbar-left'],
            'items'=>[
                ['label'=>'首页','url'=>['/index/index']],
                Yii::$app->user->isGuest?'':['label'=>'我的购物车','url'=>'/cart/index'],
                Yii::$app->user->isGuest?'':['label'=>'我的订单','url'=>'/order/index']
            ]
        ]);

    echo Nav::widget([
        'options'=>['class'=>'navbar-nav navbar-right'],
        'items'=>[
            Yii::$app->user->isGuest?['label'=>'注册','url'=>'/register/index']:['label'=>'个人中心','url'=>'user/index'],
            Yii::$app->user->isGuest?['label'=>'登录','url'=>'/site/login']:['label'=>'退出','url'=>'/site/logout']
        ]
    ]);

        NavBar::end();
    ?>

    <?php echo $content;?>
</div><!-- /.wrapper -->

<?php $this->endBody() ?>
</body>
</html>
<?php $this->endPage() ?>


```

#### 用户认证体系
1.实现认证组件User,实现认证接口\yii\web\IdentityInterface的五个方法。
2.分离前后台认证
* 修改配置文件，组件配置，
```php
'user' => [
    'identityClass' => 'app\models\User',
    'enableAutoLogin' => true,
    'idParam'=>'_user',
    'identityCookie'=>['name'=>'_user_identity','httpOnly'=>true],
    'loginUrl'=>['site/login']
],
'admin' => [
    'class'=>'yii\web\User',
    'identityClass' => 'app\models\admin',
    'enableAutoLogin' => true,
    'idParam'=>'_admin',
    'identityCookie'=>['name'=>'_admin_identity','httpOnly'=>true],
    'loginUrl'=>['admin/public/login']
],

```
* AccessControl 过滤器控制认证用户
```php
public $actions=['*'];//需要用规则去验证的方法
public $except=[];//不需要规则去验证的方法
public $mustLogin=['*'];//必须要登录的方法
/**
 * 这个装饰器方法可以抽离到公共控制器中，提高重用性
 */
public function behaviors()
{
    return [
        'access' => [
            'class' => AccessControl::class,
            'user'=>'amdin',//默认为'user'=>'user',和配置文件中的用户认证组件对应
            'only' => ['*'],
            'except'=>[],
            'rules' => [
                [
                    'allow' => true,
                    'actions' => $this->mustLogin,
                    'roles' => ['@'],
                ],
                [
                    'allow' => false,
                    'actions' => $this->mustLogin,
                    'roles' => ['?'],
                ],
                
            ],
        ],
        //控制方法的请求方式
        'verbs' => [
            'class' => VerbFilter::class,
            'actions' => [
                'logout' => ['get'],
            ],
        ],
    ];
}


```


#### RBAC

1.数据库
* auth_item #存储角色，权限
* auth_item_child #角色权限中间表
* auth_assignment #角色用户中间表
* auth_rule * 规则表

2.rbac存储方式
* 数据库的存储方式
    yii\rbac\DbManager
创建表的命令，执行这个命令之前需要先在console.php中配置authManager
```php
'authManager' => [
    'class' => 'yii\rbac\DbManager',
    'itemTable' => '{{%auth_item}}',
    'assignmentTable' => '{{%auth_assignment}}',
    'itemChildTable' => '{{%auth_item_child}}',
    'ruleTable'=>'{{%auth_rule}}',
],
    
```
```bash
 yii migrate --migrationPath=@yii/rbac/migrations
```
* 文件存储方式 
    yii\rbac\PhpManager
    @app/rbac
3.rbac相关的操作类
* yii\rbac\Item #角色后者权限的基本类，用字段type来区分
* yii\rbac\Role #Role为代表角色的类
* yii\rbac\Permission #权限类
* yii\rbac\Assignment #用户和角色的关联
* yii\rbac\Rule #角色权限的额外规则    

4.Yii2 ActiveRecord 添加额外属性（在将对象输出成json的时候会用到）
1.模型中
```php

/**
 * 添加额外的属性
 * @return array
 */
public function attributes()
{
    $attributes = parent::attributes();
    $attributes[] = 'roleText';


    return $attributes;
}

```
2.在输出json之前，将属性设值
```php
/**@var $user \backend\models\User**/
$user= $query->one();

//角色名称
$user->setAttribute('roleText', $user->getRoleText());
```