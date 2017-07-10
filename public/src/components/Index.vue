<template>
  <div class="layout-padding">
    <div class="shadow-1">
      <q-tabs
        :refs="$refs"
        default-tab="tab-1"
        class="primary shadow-1 justified"
        style="padding-top: 5px"
      >
        <q-tab name="tab-1" icon="message">My Tasks</q-tab>
        <q-tab name="tab-2" icon="today">On Going</q-tab>
        <q-tab name="tab-3" icon="assignment_turned_in">Completed</q-tab>
      </q-tabs>

      <div style="padding: 15px;">
        <div ref="tab-1">
          <h1>My Tasks</h1>
          <div class="input-group">
            <input type="text" class="form-control" placeholder="New Task" v-on:keyup.enter="createTask" v-model="newtask.name" autofocus/>
            <span class="input-group-btn">
						<button class="btn btn-primary" type="button" v-on:click="createTask">Create</button>
					</span>
          </div>
        </div>
        <div ref="tab-2">
          <h2>On Going</h2>
          <div v-for="task in tasks" style="margin-bottom: 5px;">
            <div v-if="task.done===false" class="input-group">
						<span class="input-group-btn">
							<button class="btn btn-info" type="button" :disabled="task.done===true" v-on:click="updateTask(task, true)"><i class="fa fa-check" aria-hidden="true"></i></button>
						</span>

              <input type="text" class="form-control" :disabled="task.done===true" v-model="task.name" v-on:keyup.enter="updateTask(task)"/>

              <span class="input-group-btn">
							<button class="btn btn-info" type="button" v-on:click="updateTask(task)"><i class="fa fa-edit" aria-hidden="true"></i></button>
							<button class="btn btn-danger" type="button" v-on:click="deleteTask(task.id)"><i class="fa fa-trash-o" aria-hidden="true"></i></button>
						</span>
            </div>
          </div>
        </div>
        <div ref="tab-3">
          <h2>Completed</h2>
          <div v-for="task in tasks" style="margin-bottom: 5px;">
            <div v-if="task.done===true" class="input-group">
						<span class="input-group-btn">
							<button class="btn btn-info" type="button" :disabled="task.done===true" v-on:click="updateTask(task, true)"><i class="fa fa-check" aria-hidden="true"></i></button>
						</span>

              <input type="text" class="form-control" :disabled="task.done===true" v-model="task.name" v-on:keyup.enter="updateTask(task)"/>

              <span class="input-group-btn">
							<button class="btn btn-info" type="button" :disabled="task.done===true" v-on:click="updateTask(task)"><i class="fa fa-edit" aria-hidden="true"></i></button>
							<button class="btn btn-danger" type="button" v-on:click="deleteTask(task.id)"><i class="fa fa-trash-o" aria-hidden="true"></i></button>
						</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>

  export default{
    data() {
        return {
          tasks: [],
          newtask: {}
    }
    },

    created: function() {
      this.$http.get('/tasks').then(function(res) {
        this.tasks = res.data.items ? res.data.items : [];
      });
    },

    methods: {
      createTask: function () {
        if (!$.trim(this.newtask.name)) {
          this.newtask = {};
          return;
        }
        ;

        this.newtask.done = false;

        this.$http.post('/task', this.newtask).then(function (res) {
          this.newtask.id = res.created;
          this.tasks.push(this.newtask);

          this.newtask = {};
        }),function (err) {
          console.log(err);
        };
      },

      deleteTask: function (index) {
        this.$http.delete('/task/' + index).then(function (res) {
          this.$http.get('/tasks').then(function (res) {
            this.tasks = res.data.items ? res.data.items : [];
          });
        }),function (err) {
          console.log(err);
        };
      },

      updateTask: function (task, completed) {
        if (completed) {
          task.done = true;
        }

        this.$http.put('/task', task).then(function (res) {
          this.$http.get('/tasks').then(function (res) {
            this.tasks = res.data.items ? res.data.items : [];
          });
        }),function (err) {
          console.log(err);
        };
      }
  }
  }

</script>
