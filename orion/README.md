```text

调度器
schedulerName: orion-scheduler

作用是？和schedulerName有关系么？
nodeSelector:
    gpunode: "nvidia"

下面几个环境变量的作用，和resources有何关系，是否会冲突    
 - name : ORION_GMEM
   value : "15000"
 - name : ORION_RATIO
   value : "100"
 - name: ORION_VGPU
   value: "1"
 - name: ORION_RESERVED
   value: "0"
 - name: ORION_CROSS_NODE
   value: "1"
 - name: ORION_DEVICE_TYPE
   value: "GPU"
 - name : ORION_GROUP_ID
 
 
resources:
 requests:
   virtaitech.com/gpu: 1
 limits:
   virtaitech.com/gpu: 1
 
```
