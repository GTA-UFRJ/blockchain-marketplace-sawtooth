import lmdb
import glob
size = lambda dir:sum([sum([4096 for k, v in lmdb.open(db,subdir=False).begin().cursor()]) for db in glob.glob(dir+'/*.lmdb')])
print(str(size('.')))
