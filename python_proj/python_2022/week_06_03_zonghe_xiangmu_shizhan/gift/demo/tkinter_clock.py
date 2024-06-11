import time
import tkinter as tk
from datetime import datetime


def show_time():
    time_string = datetime.now().strftime("%H:%M:%S")
    label.config(text=time_string)
    root.after(1000, show_time)


def show_time_elapse():
    t = time.time() - start_time
    min = int(t // 60)
    sec = int(t) % 60
    msec = int((t - int(t)) * 100)
    time_string = f'{min:02}:{sec:02}:{msec:02}'
    label2.config(text=time_string)
    root.after(100, show_time_elapse)


start_time = time.time()
root = tk.Tk()
root.title("Digital Clock")

# label
label = tk.Label(root, font='Arial 20', fg="yellow")
show_time()
label.pack(anchor='center')
# label2
label2 = tk.Label(root, font='Arial 20', fg="red")
show_time_elapse()
label2.pack(anchor='center')

root.mainloop()
