

    <script>

    import { onMount } from "svelte";
    import  TodoService  from "./service/todo";

      let notes = [];

      onMount(async () => {
         await TodoService.list().then(data => {
          notes = data.results;
         })
       })

      let data = {
        title: "",
        category: "",
        content: "",
        id: null
      };
      let addNote = async () => {
        const newNote = {
          title: data.title,
          description: data.description,
        };

        await TodoService.create(newNote).then((data) => {
            newNote.id = data.id
        })

        notes.unshift(newNote);
        notes = notes;
        data = {
          id: null,
          title: "",
          description: "",
        };
        console.log(notes);
      };
      let isEdit = false;
      let editNote = note => {
        isEdit = true;
        data = note;
      };
      let updateNote = () => {
        isEdit = !isEdit;
        let noteDB = {
          title: data.title,
          description: data.description,
          id: data.id
        };

        // call api
        TodoService.update(noteDB)

        let objIndex = notes.findIndex(obj => obj.id == noteDB.id);
        console.log("Before update: ", notes[objIndex]);
        notes[objIndex] = noteDB;
        data = {
          id: null,
          title: "",
          description: ""
        };
      };
      let deleteNote = async (id) => {
        let self = notes
        await TodoService.delete(id)
        await TodoService.list().then(data => {
                  self = data.results;
                 })
        notes = self
      };
    </script>
    <style>
      @import url("https://fonts.googleapis.com/css?family=Nunito&display=swap");
      * {
        font-family: "Nunito", sans-serif;
      }
    </style>
    <section>
      <div class="container">
        <div class="row mt-5 ">
          <div class="col-md-6">
            <div class="card p-2 shadow">
              <div class="card-body">
                <h5 class="card-title mb-4">Add New Note</h5>
                <form>
                  <div class="form-group">
                    <label for="title">Title</label>
                    <input
                      bind:value={data.title}
                      type="text"
                      class="form-control"
                      id="text"
                      placeholder="Note Title" />
                  </div>

                  <div class="form-group">
                    <label for="content">Description</label>
                    <textarea
                      bind:value={data.description}
                      class="form-control"
                      id="description"
                      rows="3"
                      placeholder="Note Description" />
                  </div>
                  {#if isEdit === false}
                    <button
                      type="submit"
                      on:click|preventDefault={addNote}
                      class="btn btn-primary">
                      Add Note
                    </button>
                  {:else}
                    <button
                      type="submit"
                      on:click|preventDefault={updateNote}
                      class="btn btn-info">
                      Edit Note
                    </button>
                  {/if}
                </form>
              </div>
            </div>
          </div>
          <div class="col-md-6">
            {#each notes as note}
              <div class="card mb-3">

                <div class="card-body">
                  <h5 class="card-title">{note.title}</h5>
                  <p class="card-text">{note.description}</p>
                  <button class="btn btn-info" on:click={editNote(note)}>
                    Edit
                  </button>
                  <button class="btn btn-danger" on:click={deleteNote(note.id)}>
                    Delete
                  </button>
                </div>
              </div>
            {/each}
          </div>
        </div>
      </div>
    </section>