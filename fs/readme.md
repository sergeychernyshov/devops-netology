#1

    создал простой файл 2G для тестирования
    >dd if=/dev/zero of=./sparse-file-test bs=1 count=0 seek=2G

    >ls -lash | grep sparse-file-test
    0 -rw-rw-r-- 1 vagrant vagrant 2147483648 Aug 22 18:10 sparse-file-test
    
    sparse файлы позволяют уменьшить объем больших файлов на диске.

#2

    создал жесткую ссылку на файл 
    >ln sparse-file-test sparse-file-test-link
    
    смотрю права
    >ls -lih | grep sparse-file-test
    131091 -rw-rw-r-- 2 vagrant vagrant 2.0G Aug 22 18:10 sparse-file-test
    131091 -rw-rw-r-- 2 vagrant vagrant 2.0G Aug 22 18:10 sparse-file-test-link
        
    меняю права на файл
    >chmod 777 sparse-file-test
    >ls -lih | grep sparse-file-test
    131091 -rwxrwxrwx 2 vagrant vagrant 2.0G Aug 22 18:10 sparse-file-test
    131091 -rwxrwxrwx 2 vagrant vagrant 2.0G Aug 22 18:10 sparse-file-test-link

    меняю права на ссылку
    >chmod 555 sparse-file-test-link
    >ls -lih | grep sparse-file-test
    131091 -r-xr-xr-x 2 vagrant vagrant 2.0G Aug 22 18:10 sparse-file-test
    131091 -r-xr-xr-x 2 vagrant vagrant 2.0G Aug 22 18:10 sparse-file-test-link

    меняю владельца
    >sudo chown root sparse-file-test
    >ls -lih | grep sparse-file-test
    131091 -r-xr-xr-x 2 root    vagrant 2.0G Aug 22 18:10 sparse-file-test
    131091 -r-xr-xr-x 2 root    vagrant 2.0G Aug 22 18:10 sparse-file-test-link

    меняю владельца ссылки
    >sudo chown vagrant sparse-file-test
    >ls -lih | grep sparse-file-test
    131091 -r-xr-xr-x 2 vagrant vagrant 2.0G Aug 22 18:10 sparse-file-test
    131091 -r-xr-xr-x 2 vagrant vagrant 2.0G Aug 22 18:10 sparse-file-test-link

    у файла и ссылки один владелец и одинаковые права.
    Жесткая ссылка это только ссылка на файл, которую можно разметить в любом каталоге.

#3
    
    Изменил настройки  виртуальной машины

    Vagrant.configure("2") do |config|
        config.vm.box = "bento/ubuntu-20.04" 
        config.vm.network "forwarded_port", guest: 19999, host: 19999	
        config.vm.network "forwarded_port", guest: 9100, host: 9100	
        config.vm.provider "virtualbox" do |v|
            v.memory = 4096
            v.cpus = 4
        end	
        
        config.vm.provider :virtualbox do |vb|
            lvm_experiments_disk0_path = "/tmp/lvm_experiments_disk0.vmdk"
            lvm_experiments_disk1_path = "/tmp/lvm_experiments_disk1.vmdk"
            vb.customize ['createmedium', '--filename', lvm_experiments_disk0_path, '--size', 2560]
            vb.customize ['createmedium', '--filename', lvm_experiments_disk1_path, '--size', 2560]
            vb.customize ['storageattach', :id, '--storagectl', 'SATA Controller', '--port', 1, '--device', 0, '--type', 'hdd', '--medium', lvm_experiments_disk0_path]
            vb.customize ['storageattach', :id, '--storagectl', 'SATA Controller', '--port', 2, '--device', 0, '--type', 'hdd', '--medium', lvm_experiments_disk1_path]
        end
    end
    
    настройка выполнена
    >lsblk
    NAME                 MAJ:MIN RM  SIZE RO TYPE MOUNTPOINT
    sda                    8:0    0   64G  0 disk
    ├─sda1                 8:1    0  512M  0 part /boot/efi
    ├─sda2                 8:2    0    1K  0 part
    └─sda5                 8:5    0 63.5G  0 part
      ├─vgvagrant-root   253:0    0 62.6G  0 lvm  /
      └─vgvagrant-swap_1 253:1    0  980M  0 lvm  [SWAP]
    sdb                    8:16   0  2.5G  0 disk
    sdc                    8:32   0  2.5G  0 disk

#4
    
    Запускаю интерактивный режим работы с диском
    >fdisk /dev/sdb

    проверяю разбиение диска
    > fdisk -l /dev/sdb
    Disk /dev/sdb: 2.51 GiB, 2684354560 bytes, 5242880 sectors
    Disk model: VBOX HARDDISK
    Units: sectors of 1 * 512 = 512 bytes
    Sector size (logical/physical): 512 bytes / 512 bytes
    I/O size (minimum/optimal): 512 bytes / 512 bytes
    Disklabel type: dos
    Disk identifier: 0x9576b44e
    
    Device     Boot   Start     End Sectors  Size Id Type
    /dev/sdb1          2048 4196351 4194304    2G 83 Linux
    /dev/sdb2       4196352 5242879 1046528  511M 83 Linux

#5

    Копирую партиции
    >sfdisk -d /dev/sdb|sfdisk --force /dev/sdc
    
    проверяю 
    >fdisk -l /dev/sdc
    Disk /dev/sdc: 2.51 GiB, 2684354560 bytes, 5242880 sectors
    Disk model: VBOX HARDDISK
    Units: sectors of 1 * 512 = 512 bytes
    Sector size (logical/physical): 512 bytes / 512 bytes
    I/O size (minimum/optimal): 512 bytes / 512 bytes
    Disklabel type: dos
    Disk identifier: 0x3f94c461
    
    Device     Boot   Start     End Sectors  Size Id Type
    /dev/sdc1          2048 4196351 4194304    2G 83 Linux
    /dev/sdc2       4196352 5242879 1046528  511M 83 Linux

#6

    Устанавливаю mdadm
    >apt-get install mdadm

    Пробую занулить блоки (на всякий случай, в будущем может пригодится)
    >mdadm --zero-superblock --force /dev/sd{b,c}1
    mdadm: Unrecognised md component device - /dev/sdb1
    mdadm: Unrecognised md component device - /dev/sdc1
    Блоки диска не использовались

    Удаляем старые метаданные (на всякий случай, в будущем может пригодится)
    >wipefs --all --force /dev/sd{b,c}1

    собираем RAID1
    >mdadm --create --verbose /dev/md0 -l 1 -n 2 /dev/sd{b,c}1

    >lsblk
    NAME                 MAJ:MIN RM  SIZE RO TYPE  MOUNTPOINT
    sda                    8:0    0   64G  0 disk
    ├─sda1                 8:1    0  512M  0 part  /boot/efi
    ├─sda2                 8:2    0    1K  0 part
    └─sda5                 8:5    0 63.5G  0 part
      ├─vgvagrant-root   253:0    0 62.6G  0 lvm   /
      └─vgvagrant-swap_1 253:1    0  980M  0 lvm   [SWAP]
    sdb                    8:16   0  2.5G  0 disk
    ├─sdb1                 8:17   0    2G  0 part
    │ └─md0                9:0    0    2G  0 <span style="color:red">raid1</span>
    └─sdb2                 8:18   0  511M  0 part
    sdc                    8:32   0  2.5G  0 disk
    ├─sdc1                 8:33   0    2G  0 part
    │ └─md0                9:0    0    2G  0 <span style="color:red">raid1</span>
    └─sdc2                 8:34   0  511M  0 part

    >cat /proc/mdstat
    Personalities : [linear] [multipath] [raid0] [raid1] [raid6] [raid5] [raid4] [raid10]
    md0 : active raid1 sdc1[1] sdb1[0]
          2094080 blocks super 1.2 [2/2] [UU]

#7

    собираем RAID0
    >mdadm --create --verbose /dev/md0 -l 0 -n 2 /dev/sd{b,c}2
    
    >lsblk
    NAME                 MAJ:MIN RM  SIZE RO TYPE  MOUNTPOINT
    sda                    8:0    0   64G  0 disk
    ├─sda1                 8:1    0  512M  0 part  /boot/efi
    ├─sda2                 8:2    0    1K  0 part
    └─sda5                 8:5    0 63.5G  0 part
      ├─vgvagrant-root   253:0    0 62.6G  0 lvm   /
      └─vgvagrant-swap_1 253:1    0  980M  0 lvm   [SWAP]
    sdb                    8:16   0  2.5G  0 disk
    ├─sdb1                 8:17   0    2G  0 part
    │ └─md0                9:0    0    2G  0 raid1
    └─sdb2                 8:18   0  511M  0 part
      └─md1                9:1    0 1018M  0 <span style="color:red">raid0</span>
    sdc                    8:32   0  2.5G  0 disk
    ├─sdc1                 8:33   0    2G  0 part
    │ └─md0                9:0    0    2G  0 raid1
    └─sdc2                 8:34   0  511M  0 part
      └─md1                9:1    0 1018M  0 <span style="color:red">raid0</span>
    
    >cat /proc/mdstat
    Personalities : [linear] [multipath] [raid0] [raid1] [raid6] [raid5] [raid4] [raid10]
    md1 : active raid0 sdc2[1] sdb2[0]
          1042432 blocks super 1.2 512k chunks
    
    md0 : active raid1 sdc1[1] sdb1[0]
          2094080 blocks super 1.2 [2/2] [UU]
    
#8
    Инициализируем физический раздел для использования диспетчером логических томов (LVM)
    >pvcreate /dev/md1
    Physical volume "/dev/md1" successfully created.

    >pvcreate /dev/md0
    Physical volume "/dev/md0" successfully created.

#9
    На этих физических томах создаём группу томов, которая будет называться vgRaid
    >vgcreate vgRaid /dev/md{0,1}
    Volume group "vgRaid" successfully created

    >pvs
    PV         VG        Fmt  Attr PSize    PFree
    /dev/md0   vgRaid    lvm2 a--    <2.00g   <2.00g
    /dev/md1   vgRaid    lvm2 a--  1016.00m 1016.00m 

#10

    группе томов vgRaid создаем логические тома
    >lvcreate -L 100M vgRaid /dev/md1
    Logical volume "lvol0" created.

    >lvs
    LV     VG        Attr       LSize   Pool Origin Data%  Meta%  Move Log Cpy%Sync Convert
    lvol0  vgRaid    -wi-a----- 100.00m
    root   vgvagrant -wi-ao---- <62.54g
    swap_1 vgvagrant -wi-ao---- 980.00m

#11

    Создаем файловую систему ext4
    >mkfs.ext4  /dev/vgRaid/lvol0
    mke2fs 1.45.5 (07-Jan-2020)
    Creating filesystem with 25600 4k blocks and 25600 inodes
    
    Allocating group tables: done
    Writing inode tables: done
    Creating journal (1024 blocks): done
    Writing superblocks and filesystem accounting information: done

#12
    
    Создаем директорию и монтируем наш RAID 
    >mkdir /tmp/new
    >mount /dev/vgRaid/lvol0 /tmp/new

#13

    Качаем файл из сети
    >wget https://mirror.yandex.ru/ubuntu/ls-lR.gz -O /tmp/new/test.gz

    --2021-08-23 17:29:19--  https://mirror.yandex.ru/ubuntu/ls-lR.gz
    Resolving mirror.yandex.ru (mirror.yandex.ru)... 213.180.204.183, 2a02:6b8::183
    Connecting to mirror.yandex.ru (mirror.yandex.ru)|213.180.204.183|:443... connected.
    HTTP request sent, awaiting response... 200 OK
    Length: 20980365 (20M) [application/octet-stream]
    Saving to: ‘/tmp/new/test.gz’
    
    /tmp/new/test.gz                  100%[===========================================================>]  20.01M  7.15MB/s    in 2.8s
    
    2021-08-23 17:29:22 (7.15 MB/s) - ‘/tmp/new/test.gz’ saved [20980365/20980365]

#14

    >lsblk
    NAME                 MAJ:MIN RM  SIZE RO TYPE  MOUNTPOINT
    sda                    8:0    0   64G  0 disk
    ├─sda1                 8:1    0  512M  0 part  /boot/efi
    ├─sda2                 8:2    0    1K  0 part
    └─sda5                 8:5    0 63.5G  0 part
      ├─vgvagrant-root   253:0    0 62.6G  0 lvm   /
      └─vgvagrant-swap_1 253:1    0  980M  0 lvm   [SWAP]
    sdb                    8:16   0  2.5G  0 disk
    ├─sdb1                 8:17   0    2G  0 part
    │ └─md0                9:0    0    2G  0 raid1
    └─sdb2                 8:18   0  511M  0 part
      └─md1                9:1    0 1018M  0 raid0
        └─vgRaid-lvol0   253:2    0  100M  0 lvm   /tmp/new
    sdc                    8:32   0  2.5G  0 disk
    ├─sdc1                 8:33   0    2G  0 part
    │ └─md0                9:0    0    2G  0 raid1
    └─sdc2                 8:34   0  511M  0 part
      └─md1                9:1    0 1018M  0 raid0
        └─vgRaid-lvol0   253:2    0  100M  0 lvm   /tmp/new

#15

    >gzip -t /tmp/new/test.gz
    >echo $?
    0

#16

    Мигрируем данные из RAID0 в RAID1
    >pvmove /dev/md1
    /dev/md1: Moved: 16.00%
    /dev/md1: Moved: 100.00%
    
    >lsblk
    NAME                 MAJ:MIN RM  SIZE RO TYPE  MOUNTPOINT
    sda                    8:0    0   64G  0 disk
    ├─sda1                 8:1    0  512M  0 part  /boot/efi
    ├─sda2                 8:2    0    1K  0 part
    └─sda5                 8:5    0 63.5G  0 part
      ├─vgvagrant-root   253:0    0 62.6G  0 lvm   /
      └─vgvagrant-swap_1 253:1    0  980M  0 lvm   [SWAP]
    sdb                    8:16   0  2.5G  0 disk
    ├─sdb1                 8:17   0    2G  0 part
    │ └─md0                9:0    0    2G  0 raid1
    │   └─vgRaid-lvol0   253:2    0  100M  0 lvm   /tmp/new
    └─sdb2                 8:18   0  511M  0 part
      └─md1                9:1    0 1018M  0 raid0
    sdc                    8:32   0  2.5G  0 disk
    ├─sdc1                 8:33   0    2G  0 part
    │ └─md0                9:0    0    2G  0 raid1
    │   └─vgRaid-lvol0   253:2    0  100M  0 lvm   /tmp/new
    └─sdc2                 8:34   0  511M  0 part
      └─md1                9:1    0 1018M  0 raid0

#17
    
    Переводим диск в режим сбоя    
    >mdadm /dev/md0 --fail /dev/sdb1
    mdadm: set /dev/sdb1 faulty in /dev/md0


#18

    >dmesg | grep raid1
    [  730.698003] md/raid1:md0: not clean -- starting background reconstruction
    [  730.698003] md/raid1:md0: active with 2 out of 2 mirrors
    [ 5260.643863] md/raid1:md0: Disk failure on sdb1, disabling device.
                   md/raid1:md0: Operation continuing on 1 devices.

#19
    
    >gzip -t /tmp/new/test.gz
    >echo $?
    0

#20

    >vagrant destroy
    
