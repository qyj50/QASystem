# QASystem

### 2021年秋软件工程大作业——付费问答系统

## 一、项目基本情况

### 1.1、应用背景

随着互联网技术的发展，人们能便利地从互联网获取大量信息，而信息爆炸和信息无序也增加
了人们检索、筛选、辨别信息的成本。移动支付的普及也促成了知识服务从免费过渡到付费——从
免费资源获取到“一问一答”的点对点精准问答。
提问者可以向指定回答者发起提问并支付，问题经过审核系统审核后，将会被推送给回答者，
回答者在限定的时间内进行作答。问答服务完成后将会自动对订单进行结算，回答者和平台会进行
收入和佣金分配。

### 1.2、预期用户

提问者
有支付能力且有知识付费意愿的用户，都是该平台的潜在用户
回答者
在某领域具有一定专业知识和影响力的用户，有意愿分享知识

## 二、功能需求

### 2.1、整体流程图

图 1：整体流程图

### 2.2、必选功能

付费问答系统是一个联结提问者和回答者的平台，两者通过平台进行问答沟通。用户有两种角
色：提问者和回答者。同一个用户在不同的场景中可能具有不同的角色，例如一个回答者，也可以
作为提问者向别人发起提问。
图 2：系统交互示意图
提问者可以创建问题、支付问题，回答者可接单、回答问题。在问答结束后订单自动结算：平
台抽佣、回答者获得收入。其中具体细节包括如下

**提问者**

1. 创建问答

    提问者可以在系统浏览回答者列表，向指定回答者发起提问

2. 支付

    提问者支付问答订单

3. 审核	

    问答订单支付后，问题将会进入审核列表，由审核人员审核通过后才会推送给回答
    者，若审核不通过系统会自动对提问订单进行退款

4. im 沟通

    提问者与回答者以聊天框的形式在平台上进行问答沟通，沟通内容包括文本（如有意愿可扩展至支持发送语音、视频文件），内容发送后将实时推送给对方，已结束问答服务的订单可回看聊天内容

5. 结束问答

    提问者可以主动设置问答服务完成

6. 查看订单

    提问者可在系统的订单列表页查看已完成/进行中的订单

**回答者**

1. 申请成为回答者

    注册成为系统用户后，用户可提交资料成为回答者，在资料中要描述个人简介和专业
    领域方向、设置问题定价等

2. 接单/拒单

    回答者若不想接受此问题可以主动点击拒绝接单；
    回答者若超时未接单，则订单会被取消

3. im 沟通
    回答者若未及时进行首次作答，则订单会被取消

4. 结束问答
    回答者可以主动设置问答服务完成

5. 查看订单
    回答者可在订单列表页面查看已完成/进行中的订单

6. 收入统计
    回答者可按月维度查看收入统计

**后台管理员**

- 后台管理员有 2 种角色

    - admin 管理员
        - admin 账号只有 1 个，其用户名、初始密码由系统初始化；admin 在首次登录账号后可以修改密码
        - 有且只有 admin 具有添加后台管理人员的权限，管理员的初始密码默认由系统
            初始化
        - 有且只有 admin 具有设置/更新管理员角色的权限，如 admin 添加管理员zhangsan，并设置 zhangsan 的角色为审核员
        - 有且只有 admin 具有配置系统参数的权限
            其中有哪些系统参数，举例如下：
            - 回答者问答定价设置
                生产者对问答服务定价的从价格区间，如不低于 X分，不高于 Y 分 
            - 回答者接单设置
                接单等待时长，回答者接到提问后，超过该时间未接单，问答自动取消
            - 首次作答等待时长，回答者接单后，在规定时间内没有首次作答，问答自动取消
            - 问答服务自动判别服务完成
                - 最大问答次数，如一问一答发生 X 次，则默认问答服务完成
                - 最长服务时长，如问答服务时间已超过 X 分钟后，则默认问答服务完成

    *  审核员
        * 只有具有审核员角色的管理员，才是审核员
        * 审核管理员账号仅可由 admin 添加，系统不提供管理员注册的功能
        * 审核员的初始密码由系统默认生成，审核员在首次登录后可以修改密码
        * 审核员具有问答订单审核的权限
        * 审核员具有审核回答者申请的权限（选做）
        * 管理员账户体系与用户账号体系是独立的 2 套账号体系
        * 用户账号体系具有独立的注册和登录入口
        * 管理员账号体系具有独立的登录入口
        * 一个管理员自然人也可以是一个用户（也可以申请成为回答者），且两个账户之间没有关联关系
            
            

### 2.3、可选功能

针对这两个系统，分别有一些可选功能，可按兴趣和开发能力选择性实现，具体如下：
		1）问答库

​		提问者创建问题时会增加一个“是否设置私密”选项，未设置成私密的问题则会进入问答库，能被用户搜索和浏览；设置为私密的问题将不会进入问答库，也不会被提问者和回答者以外的其他看到

​		●搜索相同问题

​	用户可搜索和浏览问答库里的问题，寻找相同或相近问题

​		●付费阅读问答内容

​	增加分享者角色，分享者可通过关注提问者、回答者或具体问题，付费阅读问答内容，其所支付的费用将会平摊给问题的提问者和回答者

​		2）评价

​		提问者在问答服务完成后可对本次问答过程以及回答者进行评价、打分，其评价分数会影响回答者在回答者列表中的排序

​		3）退款

​		提问者若对问答服务不满意，可发起退款申诉，由人工判责。申诉成功后，提问者支付的金额将会原路返回到支付账号

​		4）敏感词检查

​		付费问答审核系统能对提问中包含的敏感词进行识别，帮助审核人员提高审核效率，未通过检查的问题不能被审核通过

​		5）未读消息提醒

​		提问者、回答者订单列表对话中有未读的消息需有小红点提醒，点击查看问答对话返回后小红点消失

​		6）用户提交成为生产者时，增加审核功能

​		选做 1：用户提交成为生产者，除要求实现必选功能里的“要描述个人简介和专业领域方向、设置问题定价等”外，在流程上增加审核员审核功能：即只有审核员审核通过后才能成为生产者

​		选做 2：在选做 1 的基础上，审核员在审批提交请求时，需要用户提供额外的相关资质材料

## 三、非功能需求

### 3.1、必做部分

**问答消息的实时性、可靠性、一致性**
		提问者与回答者发送的消息要能够实时触达，不丢失、不重复，且要保证消息发送顺序与接收顺序的一致性

### 3.2、选做部分

**问答库搜索性能**
		可实现问题的模糊搜索或者是给出相似问题推荐列表，如果可以的话希望能够对搜索性能做量化评估，如召回率、准确率、相关性等；
		在搜索效率上，希望能够对万级别的文档数做到 20ms 返回搜索结果

## 四、实现要求

### 4.1、软件体系架构

付费问答系统整体流程包含浏览并选择回答者、问题创建、支付、审核、发送 im 消息等，我们可以将系统划分为用户模块、订单模块、支付模块、审核模块、im 模块。其中 im 模块的问答对话功
能，可以开发独立的 IM 即时通讯系统，也可以使用三方 SDK 包，由开发人员自行决定。

### 4.2、软件运行环境

##### 1）必做

付费问答应用主体（C 端用户和后台管理）为 Web 页面；后端环境也可由开发者选择，前端应当支持所有主流浏览器访问

##### 2）选做

付费问答 C 端用户流程可实现为 app 或小程序（付费问答后台管理仍为 Web 页面）；后端环境也可由开发者选择

### 4.3、软件开发

软件开发的具体编程语言、数据库、中间件、开发框架可由开发者选择

### 4.4、其他

订单的支付与退款，可以采用用户虚拟钱包扣减/增加余额来代替真正的金融渠道支付
