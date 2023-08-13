1、实现过程

  (1) users表新增 name 、 birthday 、 brief三个字段
  
  (2) edit接口利用session里的session_id (users.id) 更新Users表中对应id的记录（name && birthday && brief）
  
  (3) profile接口利用session里的session_id (users.id) 获取Users表中对应id的记录 （name && birthday && brief）


2、截图如下：

  （1）edit响应图：
    ![image](https://github.com/root-lzq/go_homework/assets/83163569/89c99ce8-08af-4159-b9b3-03e49bcca5b4)
  
  （2）用户名称校验图:
    ![image](https://github.com/root-lzq/go_homework/assets/83163569/bb06120e-6a52-4f0b-b568-ae0bd98be4dc)
  
  （3）日期格式校验图
    ![image](https://github.com/root-lzq/go_homework/assets/83163569/b2da2550-d289-43b8-991d-d17516fca3fa)
  
  （4）简介校验图：
    ![image](https://github.com/root-lzq/go_homework/assets/83163569/8bd5431a-615a-4720-a43c-0ba7ae77d361)
  
  （5）profile响应图
    ![image](https://github.com/root-lzq/go_homework/assets/83163569/44b7ff4e-39ae-4140-af0d-274433c370fd)
