#include <stdio.h>

struct ListNode
{
    int val;
    struct ListNode *next;
};

struct ListNode *getIntersectionNode(struct ListNode *headA, struct ListNode *headB)
{
    if (headA == NULL || headB == NULL)
    {
        return NULL;
    }
    struct ListNode *pa = headA;
    struct ListNode *pb = headB;
    int lena = 0;
    int lenb = 0;
    while (pa != NULL)
    {
        lena++;
        pa = pa->next;
    }
    while (pb != NULL)
    {
        lenb++;
        pb = pb->next;
    }
    int t = lena - lenb;
    pa = headA;
    pb = headB;
    if (t > 0)
    {
        while (t > 0)
        {
            pa = pa->next;
            t--;
        }
    }
    else
    {
        while (t < 0)
        {
            pb = pb->next;
            t++;
        }
    }
    do
    {
        if (pa == pb)
        {
            return pa;
        }
        pa = pa->next;
        pb = pb->next;
    } while (pa != NULL);
    return NULL;
}