# Storage

## Contents

- Introduction(简介)
- Functionality(功能)
- Data organization and representation
- Hierarchy of storage
  - Primary storage
  - Secondary storage
  - Tertiary storage
  - Off-line storage
- Characteristics of storage
  - Volatility
  - Mutability
  - Accessibility
  - Addressability
  - Capacity
  - Performance
  - Energy use
  - Security

## Introduction

**Computer data storage**, often called **storage** or **memory**, is a technology consisting of computer components and recording media that are used to retain digital data. It is a core function and fundamental component of computers. (存储是计算机的核心功能和重要组成部分)

The **central processing unit** (CPU) of a computer is what manipulates data by performing computations. In practice, almost all computers use a storage hierarchy which puts fast but expensive and small storage options close to the CPU and slower but larger and cheaper options farther away. Generally the fast volatile technologies (which lose data when off power) are referred to as "memory", while slower persistent technologies are referred to as "storage". (CPU通过计算操纵数据。实际中，计算机采用分层存储，靠近CPU、运行速度较快的部分被称为“memory”，反之称为“storage”)

## Functionality

Without a significant amount of memory, a computer would merely be able to perform fixed operations and immediately output the result. It would have to be reconfigured to change its behavior. This is acceptable for devices such as desk calculators, digital signal processors, and other specialized devices. Von Neumann machines differ in having a memory in which they store their operating instructions and data. Such computers are more versatile(灵活的) in that they do not need to have their hardware reconfigured for each new program, but can simply be reprogrammed with new in-memory instructions; they also tend to be simpler to design, in that a relatively simple processor may keep state between successive computations to build up complex procedural results. Most modern computers are von Neumann machines.(没有存储，计算机只能够执行固定的操作并且立即输出结果。拥有存储技术的计算机存在以下两条优点：1.灵活性。每一个新的程序可以通过内在的指令进行编程，而不是靠硬件重新编程。2.简易性。处理器保存着连续的计算的状态，最终产生复杂的输出结果。)

## Data organization and representation

A modern digital computer represents data using the binary numeral system. Text, numbers, pictures, audio, and nearly any other form of information can be converted into a string of bits, or binary digits, each of which has a value of 1 or 0. The most common unit of storage is the byte, equal to 8 bits. A piece of information can be handled by any computer or device whose storage space is large enough to accommodate *the binary representation of the piece of information*, or simply data. For example, the complete works of Shakespeare, about 1250 pages in print, can be stored in about five megabytes(40 million bits) with one byte per character.(现代计算机使用二进制存储数据)

Data are encoded by assigning a bit pattern to each character, digit, or multimedia object. Many standards exist for encoding (e.g., character encodings like ASCII, image encodings like JPEG, video encodings like MPEG-4.(不同类型的数据往往存在着不同的编码标准)

By adding bits to each encoded unit, redundancy allows the computer to both detect errors in coded data and correct them based on mathematical algorithms. Errors generally occur in low probabilities due to random bit value flipping, or "physical bit fatigue", loss of the physical bit in storage of its ability to maintain a distinguishable value (0 or 1), or due to errors in inter or intra-computer communication. A random bit flip(e.g., due to random radiation) is typically corrected upon detection. A bit, or a group of malfunctioning physical bits (not always the specific defective bit is known; group definition depends on specific storage device) is typically automatically fenced-out, taken out of use by the device, and replaced with another functioning equivalent group in the device, where the corrected bit values are restored (if possible). The cyclic redundancy check(CRC) method is typically used in communications and storage for error detection. A detected error is then retried.(在进行数据编码编码的时候，往往会增添一校验位来检测是否出现错误)

Data compression methods allow in many cases (such as a database) to represent a string of bits by a shorter bit string ("compress") and reconstruct the original string ("decompress") when needed. This utilizes substantially less storage (tens of percents) for many types of data at the cost of more computation (compress and decompress when needed). Analysis of trade-off between storage cost saving and costs of related computations and possible delays in data availability is done before deciding whether to keep certain data compressed or not.(在数据存储的时候往往还要考虑数据压缩)

For security reasons certain types of data (e.g., credit-card information) may be kept encrypted in storage to prevent the possibility of unauthorized information reconstruction from chunks of storage snapshots.(为了安全考虑有些数据存储时还要考虑加密)

>总结：
>- 现代计算机使用二进制进行数据存储。
>- 不同类型的数据往往存在着不同的编码标准。
>- 在进行数据存储时，通常需要考虑：
>	- 校验位
>	- 压缩
>	- 加密

## Hierarchy of storage

[![img](https://upload.wikimedia.org/wikipedia/commons/thumb/3/3e/Computer_storage_types.svg/350px-Computer_storage_types.svg.png)](https://en.wikipedia.org/wiki/File:Computer_storage_types.svg)

Generally, the lower a storage is in the hierarchy, the lesser its bandwidth and the greater its access latency is from the CPU. This traditional division of storage to primary, secondary, tertiary and off-line storage is also guided by cost per bit.

In contemporary usage, "memory" is usually semiconductor storage read-write random-access memory, typically DRAM (dynamic RAM) or other forms of fast but temporary storage. "Storage" consists of storage devices and their media not directly accessible by the CPU, typically hard disk drives, optical disc drives, and other devices slower than RAM but non-volatile (retaining contents when powered down).

Historically, **memory** has been called **core memory**, **main memory**, **real storage** or **internal memory**. Meanwhile, non-volatile storage devices have been referred to as **secondary storage**, **external memory** or **auxiliary/peripheral storage**.

### Primary storage

**Primary storage** (also known as *main memory*, *internal memory* or *prime memory*), often referred to simply as *memory*, is the only one directly accessible to the CPU. The CPU continuously reads instructions stored there and executes them as required. Any data actively operated on is also stored there in uniform manner.

Historically, early computers used delay lines, Williams tubes, or rotating magnetic drums as primary storage. By 1954, those unreliable methods were mostly replaced by magnetic core memory. Core memory remained dominant until the 1970s, when advances in integrated circuittechnology allowed semiconductor memory to become economically competitive.

This led to modern random-access memory(RAM). It is small-sized, light, but quite expensive at the same time. (The particular types of RAM used for primary storage are also volatile, i.e. they lose the information when not powered).

As shown in the diagram, traditionally there are two more sub-layers of the primary storage, besides main large-capacity RAM:

- Processor registers are located inside the processor. Each register typically holds a word of data (often 32 or 64 bits). CPU instructions instruct the arithmetic logic unit to perform various calculations or other operations on this data (or with the help of it). Registers are the fastest of all forms of computer data storage.
- Processor cache is an intermediate stage between ultra-fast registers and much slower main memory. It was introduced solely to improve the performance of computers. Most actively used information in the main memory is just duplicated in the cache memory, which is faster, but of much lesser capacity. On the other hand, main memory is much slower, but has a much greater storage capacity than processor registers. Multi-level hierarchical cache setup is also commonly used—*primary cache* being smallest, fastest and located inside the processor; *secondary cache* being somewhat larger and slower.

Main memory is directly or indirectly connected to the central processing unit via a *memory bus*. It is actually two buses (not on the diagram): an address bus and a data bus. The CPU firstly sends a number through an address bus, a number called memory address, that indicates the desired location of data. Then it reads or writes the data in the memory cellsusing the data bus. Additionally, a memory management unit(MMU) is a small device between CPU and RAM recalculating the actual memory address, for example to provide an abstraction of virtual memory or other tasks.

As the RAM types used for primary storage are volatile (uninitialized at start up), a computer containing only such storage would not have a source to read instructions from, in order to start the computer. Hence, non-volatile primary storage containing a small startup program is used to bootstrap the computer, that is, to read a larger program from non-volatile **secondary** storage to RAM and start to execute it. A non-volatile technology used for this purpose is called ROM, for read-only memory(the terminology may be somewhat confusing as most ROM types are also capable of **random access**).

Many types of "ROM" are not literally *read only*, as updates to them are possible; however it is slow and memory must be erased in large portions before it can be re-written. Some embedded systems run programs directly from ROM (or similar), because such programs are rarely changed. Standard computers do not store non-rudimentary programs in ROM, and rather, use large capacities of secondary storage, which is non-volatile as well, and not as costly.

Recently, **primary storage** and **secondary storage** in some uses refer to what was historically called, respectively, **secondary storage** and **tertiary storage**.

### Secondary storage

[![img](https://upload.wikimedia.org/wikipedia/commons/thumb/9/97/35-Desktop-Hard-Drive.jpg/220px-35-Desktop-Hard-Drive.jpg)](https://en.wikipedia.org/wiki/File:35-Desktop-Hard-Drive.jpg)

**Secondary storage** (also known as **external memory** or **auxiliary storage**), differs from primary storage in that it is not directly accessible by the CPU. The computer usually uses its input/output channels to access secondary storage and transfer the desired data to primary storage. Secondary storage is non-volatile (retaining data when power is shut off). Modern computer systems typically have two orders of magnitude more secondary storage than primary storage because secondary storage is less expensive.

In modern computers, hard disk drives(HDDs) or solid-state drives(SSDs) are usually used as secondary storage. The access time per byte for HDDs or SSDs is typically measured in milliseconds(one thousandth seconds), while the access time per byte for primary storage is measured in nanoseconds(one billionth seconds). Thus, secondary storage is significantly slower than primary storage. Rotating optical storage devices, such as CDand DVDdrives, have even longer access times. Other examples of secondary storage technologies include USB flash drives, floppy disks, magnetic tape, paper tape, punched cards, and RAM disks.

Once the disk read/write head on HDDs reaches the proper placement and the data, subsequent data on the track are very fast to access. To reduce the seek time and rotational latency, data are transferred to and from disks in large contiguous blocks. Sequential or block access on disks is orders of magnitude faster than random access, and many sophisticated paradigms have been developed to design efficient algorithms based upon sequential and block access. Another way to reduce the I/O bottleneck is to use multiple disks in parallel in order to increase the bandwidth between primary and secondary memory.

Secondary storage is often formatted according to a file system format, which provides the abstraction necessary to organize data into files and directories, while also providing metadata describing the owner of a certain file, the access time, the access permissions, and other information.

Most computer operating systems use the concept of virtual memory, allowing utilization of more primary storage capacity than is physically available in the system. As the primary memory fills up, the system moves the least-used chunks to a swap file or page file on secondary storage, retrieving them later when needed. If a lot of pages are moved to slower secondary storage, the system performance is degraded.

### Tertiary storage
[![img](https://upload.wikimedia.org/wikipedia/commons/thumb/c/c9/StorageTek_Powderhorn_tape_library.jpg/220px-StorageTek_Powderhorn_tape_library.jpg)](https://en.wikipedia.org/wiki/File:StorageTek_Powderhorn_tape_library.jpg)



A large , with tape cartridges placed on shelves in the front, and a robotic arm moving in the back. Visible height of the library is about 180 cm.

**Tertiary storage** or **tertiary memory** is a level below secondary storage. Typically, it involves a robotic mechanism which will *mount* (insert) and *dismount* removable mass storage media into a storage device according to the system's demands; such data are often copied to secondary storage before use. It is primarily used for archiving rarely accessed information since it is much slower than secondary storage (e.g. 5–60 seconds vs. 1–10 milliseconds). This is primarily useful for extraordinarily large data stores, accessed without human operators. Typical examples include tape libraries and optical jukeboxes.

When a computer needs to read information from the tertiary storage, it will first consult a catalog databaseto determine which tape or disc contains the information. Next, the computer will instruct a robotic arm to fetch the medium and place it in a drive. When the computer has finished reading the information, the robotic arm will return the medium to its place in the library.

Tertiary storage is also known as *nearline storage* because it is "near to online". The formal distinction between online, nearline, and offline storage is:

- Online storage is immediately available for I/O.
- Nearline storage is not immediately available, but can be made online quickly without human intervention.
- Offline storage is not immediately available, and requires some human intervention to become online.

For example, always-on spinning hard disk drives are online storage, while spinning drives that spin down automatically, such as in massive arrays of idle disks, are nearline storage. Removable media such as tape cartridges that can be automatically loaded, as in tape libraries, are nearline storage, while tape cartridges that must be manually loaded are offline storage.

### Off-line storage

*Off-line storage* is a computer data storage on a medium or a device that is not under the control of a processing unit. The medium is recorded, usually in a secondary or tertiary storage device, and then physically removed or disconnected. It must be inserted or connected by a human operator before a computer can access it again. Unlike tertiary storage, it cannot be accessed without human interaction.

Off-line storage is used to transfer information, since the detached medium can be easily physically transported. Additionally, in case a disaster, for example a fire, destroys the original data, a medium in a remote location will probably be unaffected, enabling disaster recovery. Off-line storage increases general information security, since it is physically inaccessible from a computer, and data confidentiality or integrity cannot be affected by computer-based attack techniques. Also, if the information stored for archival purposes is rarely accessed, off-line storage is less expensive than tertiary storage.

In modern personal computers, most secondary and tertiary storage media are also used for off-line storage. Optical discs and flash memory devices are most popular, and to much lesser extent removable hard disk drives. In enterprise uses, magnetic tape is predominant. Older examples are floppy disks, Zip disks, or punched cards.

## Characteristics of storage

Storage technologies at all levels of the storage hierarchy can be differentiated by evaluating certain core characteristics as well as measuring characteristics specific to a particular implementation. These core characteristics are volatility, mutability, accessibility, and addressability. For any particular implementation of any storage technology, the characteristics worth measuring are capacity and performance.

### Volatility

Non-volatile memory retains the stored information even if not constantly supplied with electric power. It is suitable for long-term storage of information. Volatile memory requires constant power to maintain the stored information. The fastest memory technologies are volatile ones, although that is not a universal rule. Since the primary storage is required to be very fast, it predominantly uses volatile memory.(非易变内存存储着持久化的信息，易变内存需要持久的供电)

Dynamic random-access memory is a form of volatile memory that also requires the stored information to be periodically reread and rewritten, or refreshed, otherwise it would vanish. Static random-access memory is a form of volatile memory similar to DRAM with the exception that it never needs to be refreshed as long as power is applied; it loses its content when the power supply is lost.(SRAM和DRAM在保留信息时间和是否需要供电方面不相同)

An uninterruptible power supply (UPS) can be used to give a computer a brief window of time to move information from primary volatile storage into non-volatile storage before the batteries are exhausted. Some systems, for example EMC Symmetrix, have integrated batteries that maintain volatile storage for several minutes.

### Mutability

- Read/write storage or mutable storage 

  Allows information to be overwritten at any time. A computer without some amount of read/write storage for primary storage purposes would be useless for many tasks. Modern computers typically use read/write storage also for secondary storage.

- Slow write, fast read storage 

  Read/write storage which allows information to be overwritten multiple times, but with the write operation being much slower than the read operation.

- Write once storage 

  Write Once Read Many (WORM) allows the information to be written only once at some point after manufacture. Examples include semiconductor programmable read-only memory and CD-R.

- Read only storage 

  Retains the information stored at the time of manufacture. Examples include mask ROM ICs and CD-ROM.

>总结：
>常见的存储设备的可变性特点：
>- 可读可写
>- 慢写快读
>- write once
>- read only

### Accessibility

- Random access

  Any location in storage can be accessed at any moment in approximately the same amount of time. Such characteristic is well suited for primary and secondary storage. Most semiconductor memories and disk drives provide random access.

- Sequential access

  The accessing of pieces of information will be in a serial order, one after the other; therefore the time to access a particular piece of information depends upon which piece of information was last accessed. Such characteristic is typical of off-line storage.
  
>总结：
>两种数据获取方式：
>- random access
>- sequential access

### Addressability

- Location-addressable 

  Each individually accessible unit of information in storage is selected with its numerical memory address. In modern computers, location-addressable storage usually limits to primary storage, accessed internally by computer programs, since location-addressability is very efficient, but burdensome for humans.

- File addressable

  Information is divided into **files** of variable length, and a particular file is selected with human-readable directory and file names. The underlying device is still location-addressable, but the operating system of a computer provides the file system abstraction to make the operation more understandable. In modern computers, secondary, tertiary and off-line storage use file systems.

- Content-addressable

  Each individually accessible unit of information is selected based on the basis of (part of) the contents stored there. Content-addressable storage can be implemented using software (computer program) or hardware (computer device), with hardware being faster but more expensive option. Hardware content addressable memory is often used in a computer's CPU cache.

### Capacity

- Raw capacity

  The total amount of stored information that a storage device or medium can hold. It is expressed as a quantity of bits or bytes.

- Memory storage density

  The compactness of stored information. It is the storage capacity of a medium divided with a unit of length, area or volume (e.g. 1.2 megabytes per square inch).

### Performance

- Latency

  The time it takes to access a particular location in storage. The relevant unit of measurement is typically nanosecond for primary storage, millisecond for secondary storage, and second for tertiary storage. It may make sense to separate read latency and write latency (especially for non-volatile memory and in case of sequential access storage, minimum, maximum and average latency.

- Throughput

  The rate at which information can be read from or written to the storage. In computer data storage, throughput is usually expressed in terms of megabytes per second (MB/s), though bit rate may also be used. As with latency, read rate and write rate may need to be differentiated. Also accessing media sequentially, as opposed to randomly, typically yields maximum throughput.

- Granularity

  The size of the largest "chunk" of data that can be efficiently accessed as a single unit, e.g. without introducing additional latency.

- Reliability

  The probability of spontaneous bit value change under various conditions, or overall failure rate.

### Energy use

- Storage devices that reduce fan usage, automatically shut-down during inactivity, and low power hard drives can reduce energy consumption by 90 percent.
- 2.5-inch hard disk drives often consume less power than larger ones. Low capacity  have no moving parts and consume less power than hard disks. Also, memory may use more power than hard disks.Large caches, which are used to avoid hitting the , may also consume a large amount of power.

### Security

Full disk encryption, volume and virtual disk encryption, andor file/folder encryption is readily available for most storage devices.

Hardware memory encryption is available in Intel Architecture, supporting Total Memory Encryption (TME) and page granular memory encryption with multiple keys (MKTME). and in M7 generation since October 2015.


<i>Reference:</i>

[1] [Wiki: Computer data storage](https://en.wikipedia.org/wiki/Computer_data_storage)