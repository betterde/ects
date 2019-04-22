<template>
  <div class="main-content">
    <div class="panel">
      <div class="panel-header">
        <div class="panel-tools">
          <el-row :gutter="20">
            <el-col :span="16">
              <el-button type="primary" plain>Create</el-button>
            </el-col>
            <el-col :span="8">
              <el-input placeholder="Search in here" v-model="params.search"><i slot="prefix" class="el-input__icon el-icon-search"></i></el-input>
            </el-col>
          </el-row>
        </div>
      </div>
      <div class="panel-body">
        <el-table :data="users" style="width: 100%">
          <el-table-column type="expand">
            <template slot-scope="props">
              <el-form label-position="top" inline class="table-expand">
                <el-row :gutter="10">
                  <el-col :span="12">
                    <el-form-item label="ID">
                      <span>{{ props.row.id }}</span>
                    </el-form-item>
                  </el-col>
                  <el-col :span="12">
                    <el-form-item label="Team">
                      <span>{{ props.row.team_id === "" ? "Not on any team" : props.row.team_id }}</span>
                    </el-form-item>
                  </el-col>
                </el-row>
                <el-row :gutter="10">
                  <el-col :span="12">
                    <el-form-item label="Role">
                      <span>{{ props.row.role_id === "" ? "Unspecified role" : props.row.role_id}}</span>
                    </el-form-item>
                  </el-col>
                  <el-col :span="12">
                    <el-form-item label="Create at">
                      <span>{{ props.row.created_at }}</span>
                    </el-form-item>
                  </el-col>
                </el-row>
              </el-form>
            </template>
          </el-table-column>
          <el-table-column prop="name" label="Name" width="200"></el-table-column>
          <el-table-column prop="email" label="Email" width="300"></el-table-column>
          <el-table-column label="Manager" width="100">
            <template slot-scope="scope">
              <el-tag v-if="scope.row.manager === 0" size="medium" type="info">NO</el-tag>
              <el-tag v-else size="medium">Yes</el-tag>
            </template>
          </el-table-column>
          <el-table-column label="Status">
            <template slot-scope="scope">
              <el-tag v-if="scope.row.deleted_at === ''" size="medium">Normal</el-tag>
              <el-tag v-else size="medium" type="info">Disabled</el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="option" label="Action" width="130">
            <template slot-scope="scope">
              <el-button size="mini" icon="el-icon-edit" circle @click="handleEdit(scope.$index, scope.row)"></el-button>
              <el-button size="mini" icon="el-icon-tickets" plain circle @click="handleDelete(scope.$index, scope.row)"></el-button>
              <el-button size="mini" icon="el-icon-delete" type="danger" plain circle @click="handleDelete(scope.$index, scope.row)"></el-button>
            </template>
          </el-table-column>
        </el-table>
        <div class="pagination">
          <el-pagination background layout="prev, pager, next" :total="11"></el-pagination>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import api from '../apis'

  export default {
    name: "User",
    data() {
      return {
        loading: false,
        params: {
          search: ""
        },
        users: [],
        meta: {
          limit: 10,
          page: 1,
          total: 0
        }
      }
    },
    methods: {
      handleEdit(index, row) {
        window.console.log(index, row);
      },
      handleDelete(index, row) {
        window.console.log(index, row);
      },
      fetchUsers() {
        this.loading = true;
        api.user.fetch(this.params).then(res => {
          this.users = res.data;
          this.meta = res.meta;
        }).catch(err => {
          this.$message.warning(err.message)
        });
        this.loading = false;
      }
    },
    mounted() {
      this.fetchUsers()
    }
  }
</script>

<style lang="scss">

</style>
