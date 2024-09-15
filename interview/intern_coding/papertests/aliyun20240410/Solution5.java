package interview.papertests.aliyun20240410;

import java.util.*;

public class Solution5 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int T = in.nextInt();
        for(int i = 0;i<T;i++){
            int n = in.nextInt();
            int k = in.nextInt();
            Map<Integer,Integer> map = new HashMap<>();
            for(int j = 0;j<n;j++){
                int x = in.nextInt();
                map.put(x, map.getOrDefault(x,0)+1);
            }
            List<Integer> counts = new ArrayList<>();
            counts.addAll(map.values());
            Collections.sort(counts);

            int[] s = new int[counts.size()];   //前缀和
            s[0] = counts.get(0);
            for(int j = 1;j<s.length;j++){
                s[j] = s[j-1] + counts.get(j);
            }
            int res = Integer.MIN_VALUE;
//            for (int j = 0;j< counts.size();j++){
//                System.out.print(counts.get(j)+" ");
//            }
//            System.out.println();
//            for (int j = 0;j< counts.size();j++){
//                System.out.print(s[j]+" ");
//            }
//            System.out.println();

            for(Map.Entry<Integer,Integer> nC:map.entrySet()){
                int num = nC.getKey();
                int value = nC.getValue();
                int index = Collections.binarySearch(counts, value);
                int x = index >= 1 ?s[index-1]+value+(value-1)*(counts.size()-1-index):value+(value-1)*(counts.size()-1-index);
                //System.out.println("num = " + num + " , count = " + value + " shu = " + x);
                if(x >=k){
                    res = Math.max(res,num);
                }
            }
            System.out.println(res);
        }
    }
}
