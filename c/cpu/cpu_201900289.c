#include <linux/init.h>
#include <linux/sched/signal.h>
#include <linux/proc_fs.h>
#include <linux/seq_file.h>
#include <linux/module.h>
#include <linux/kernel.h>
#include <linux/sched.h>
#include <linux/version.h>
#include <linux/fs.h>

MODULE_LICENSE("GPL");
MODULE_AUTHOR("Juan Pablo Rojas Chinchilla");
MODULE_DESCRIPTION("modulo de cpu");

#if LINUX_VERSION_CODE >= KERNEL_VERSION(5,6,0)
#define HAVE_PROC_OPS
#endif

struct task_struct *task;
struct task_struct *taskChild;
struct list_head *list;

static int proc_cpu(struct seq_file * file, void *v){
    int running = 0;
    int sleeping =0;
    int zombie = 0;
    int stopped = 0;

    seq_printf(file,"{\n\"processes\":[\n");

    for_each_process(task){
        seq_printf(file, "{");
        seq_printf(file,"\"pid\":%d,\n",task->pid);
        seq_printf(file,"\"name\":\"%s\",\n",task->comm);
        seq_printf(file, "\"user\": %d,\n",task->cred->uid);
        seq_printf(file,"\"state\":%ld,\n",task->__state);

        seq_printf(file,"\"child\":[\n");
        list_for_each(list, &task->children){
            taskChild = list_entry(list,struct task_struct,sibling);
            seq_printf(file,"%d,",taskChild->pid);
        }
        seq_printf(file,"\n]");

        if(task->__state == 0){
            running +=1;
        }
        else if (task->__state ==1){
            sleeping +=1;
        }
        else if (task->__state == 4){
            zombie +=1;
        }
        else if (task->__state == 5){
            stopped +=1;
        }
        seq_printf(file, "},\n");
    }

    seq_printf(file, "],\n");
    seq_printf(file, "\"running\":%d,\n",running);
    seq_printf(file, "\"sleeping\":%d,\n",sleeping);
    seq_printf(file, "\"zombie\":%d,\n",zombie);
    seq_printf(file, "\"stopped\":%d,\n",stopped);
    seq_printf(file, "\"total\":%d\n",running+sleeping+zombie+stopped);
    seq_printf(file,"}\n");
    return 0;

}

static int open_cpu(struct inode *inode, struct file * file){
    return single_open(file,proc_cpu,NULL);
}

#ifdef HAVE_PROC_OPS
static const struct proc_ops operations = {
  .proc_open = open_cpu,
  .proc_read = seq_read,
  .proc_lseek = seq_lseek,
  .proc_release = single_release,
};
#else
static const struct file_operations operations = {
  .owner = THIS_MODULE,
  .open = open_cpu,
  .read = seq_read,
  .llseek = seq_lseek,
  .release = single_release,
};
#endif

static int start(void){
    proc_create("cpu_201900289",0,NULL,&operations);
    printk(KERN_INFO "Cargando modulo de cpu\n");
    printk(KERN_INFO "Nombre: Juan Pablo Rojas Chinchilla\n");
    return 0;
}

static void __exit finish(void){
    remove_proc_entry("cpu_201900289",NULL);
    printk(KERN_INFO "Removiendo modulo cpu\n");
    printk(KERN_INFO "Diciembre 2021\n");
}

module_init(start);
module_exit(finish);