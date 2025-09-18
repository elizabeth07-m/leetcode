type TaskManager struct {
	pipeline *TaskList
	tasks    map[int]*Task
}

type Task struct {
	taskId   int
	userId   int
	priority int
	index    int
}

type TaskList []*Task

func (t TaskList) Len() int { return len(t) }

func (t TaskList) Less(i, j int) bool {
	if t[i].priority != t[j].priority {
		return t[i].priority > t[j].priority
	}
	return t[i].taskId > t[j].taskId
}

func (t TaskList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
	t[i].index = i
	t[j].index = j
}

func (t *TaskList) Push(x any) {
	task := x.(*Task)
	task.index = len(*t)
	*t = append(*t, task)
}

func (t *TaskList) Pop() any {
	old := *t
	n := len(old)
	x := old[n-1]
	old[n-1] = nil
	*t = old[:n-1]
	return x
}

func Constructor(tasks [][]int) TaskManager {
	tl := &TaskList{}
	m := make(map[int]*Task)
	for _, task := range tasks {
		t := &Task{
			userId:   task[0],
			taskId:   task[1],
			priority: task[2],
		}
		m[task[1]] = t
		heap.Push(tl, t)
	}

	return TaskManager{
		tasks:    m,
		pipeline: tl,
	}
}

func (this *TaskManager) Add(userId int, taskId int, priority int) {
	t := &Task{
		userId:   userId,
		taskId:   taskId,
		priority: priority,
	}
	this.tasks[taskId] = t
	heap.Push(this.pipeline, t)
}

func (this *TaskManager) Edit(taskId int, newPriority int) {
	t := this.tasks[taskId]
	t.priority = newPriority
	heap.Fix(this.pipeline, t.index)
}

func (this *TaskManager) Rmv(taskId int) {
	t := this.tasks[taskId]
	heap.Remove(this.pipeline, t.index)
	delete(this.tasks, taskId)
}

func (this *TaskManager) ExecTop() int {
	if this.pipeline.Len() > 0 {
		t := heap.Pop(this.pipeline).(*Task)
		delete(this.tasks, t.taskId)
		return t.userId
	}
	return -1
}

/**
 * Your TaskManager object will be instantiated and called as such:
 * obj := Constructor(tasks);
 * obj.Add(userId,taskId,priority);
 * obj.Edit(taskId,newPriority);
 * obj.Rmv(taskId);
 * param_4 := obj.ExecTop();
 */