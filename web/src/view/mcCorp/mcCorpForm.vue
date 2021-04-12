<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="企业外部联系人secret:">
                <el-input v-model="formData.contactSecret" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="企业通讯录secret:">
                <el-input v-model="formData.employeeSecret" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="回调消息加密串:">
                <el-input v-model="formData.encodingAesKey" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="事件回调地址:">
                <el-input v-model="formData.eventCallback" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="企业名称:">
                <el-input v-model="formData.name" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="企业代码(企业统一社会信用代码):">
                <el-input v-model="formData.socialCode" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="租户ID:"><el-input v-model.number="formData.tenantId" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="回调token:">
                <el-input v-model="formData.token" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="企业微信ID:">
                <el-input v-model="formData.wxCorpid" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           <el-form-item>
           <el-button @click="save" type="primary">保存</el-button>
           <el-button @click="back" type="primary">返回</el-button>
           </el-form-item>
    </el-form>
</div>
</template>

<script>
import {
    createMcCorp,
    updateMcCorp,
    findMcCorp
} from "@/api/mcCorp";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "McCorp",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
            contactSecret:"",
            employeeSecret:"",
            encodingAesKey:"",
            eventCallback:"",
            name:"",
            socialCode:"",
            tenantId:0,
            token:"",
            wxCorpid:"",
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createMcCorp(this.formData);
          break;
        case "update":
          res = await updateMcCorp(this.formData);
          break;
        default:
          res = await createMcCorp(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功"
        })
      }
    },
    back(){
        this.$router.go(-1)
    }
  },
  async created() {
   // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if(this.$route.query.id){
    const res = await findMcCorp({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.remcCorp
       this.type = "update"
     }
    }else{
     this.type = "create"
   }
  
}
};
</script>

<style>
</style>